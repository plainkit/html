package spec

import (
	"fmt"
	"sort"
	"strings"

	"github.com/delaneyj/gostar/cfg"
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

// Loader handles loading and parsing of SVG specification from gostar
type Loader struct {
	svgSpec *pb.Namespace
}

// NewLoader creates a new spec loader using gostar SVG data
func NewLoader() *Loader {
	return &Loader{svgSpec: cfg.SVG}
}

// getGostarSpec returns the gostar SVG specification
func (l *Loader) getGostarSpec() *pb.Namespace {
	return l.svgSpec
}

// isVoidElement checks if an SVG element is self-closing based on gostar data
func (l *Loader) isVoidElement(tagName string) bool {
	for _, element := range l.svgSpec.Elements {
		if element.Tag == tagName {
			return element.NoChildren
		}
	}
	return false
}

// isAttributeBoolean determines if an attribute should be treated as boolean
func (l *Loader) isAttributeBoolean(attr *pb.Attribute) bool {
	if attr.Type == nil {
		return false
	}
	// Check if the attribute type contains "bool:true" in its string representation
	typeStr := fmt.Sprintf("%+v", attr.Type)
	return strings.Contains(typeStr, "bool:true")
}

// LoadAllTagSpecs loads all SVG tag specifications from gostar data
func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	fmt.Println("Loading SVG specifications from gostar...")
	specs := l.convertGostarToTagSpecs()
	fmt.Printf("Loaded %d SVG element specifications\n", len(specs))
	return specs, nil
}

// LoadGlobalAttributes loads SVG global attributes from gostar data
func (l *Loader) LoadGlobalAttributes() ([]Attribute, error) {
	fmt.Println("Loading SVG global attributes from gostar...")
	globalAttrs := l.extractGlobalAttributes()

	var attributes []Attribute
	for attrName, attrType := range globalAttrs {
		if attrName == "" {
			continue
		}

		field := camelCase(attrName)
		attr := Attribute{
			Field: field,
			Attr:  attrName,
			Type:  attrType,
		}

		attributes = append(attributes, attr)
	}

	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Attr < attributes[j].Attr
	})

	return attributes, nil
}

// CollectAllAttributes collects unique attributes from all SVG tag specs
func (l *Loader) CollectAllAttributes(specs []TagSpec) map[string]Attribute {
	allAttributes := make(map[string]Attribute)

	for _, spec := range specs {
		for _, attr := range spec.Attributes {
			key := strings.ToLower(attr.Attr)
			if existing, exists := allAttributes[key]; exists {
				// If we have a conflict, prefer hyphenated version (produces better CamelCase)
				// or string over bool for flexibility
				if strings.Contains(attr.Attr, "-") && !strings.Contains(existing.Attr, "-") {
					allAttributes[key] = attr
				} else if existing.Type == "bool" && attr.Type == "string" {
					allAttributes[key] = attr
				}
			} else {
				allAttributes[key] = attr
			}
		}
	}

	return allAttributes
}

func (l *Loader) convertGostarToTagSpecs() []TagSpec {
	var specs []TagSpec

	globalAttrs := l.extractGlobalAttributes()

	// SVG elements that should always be void/self-closing
	svgVoidElements := map[string]bool{
		"path":             true,
		"circle":           true,
		"ellipse":          true,
		"line":             true,
		"rect":             true,
		"polygon":          true,
		"polyline":         true,
		"use":              true,
		"stop":             true,
		"animate":          true,
		"animateMotion":    true,
		"animateTransform": true,
	}

	// Get HTML elements to skip from HTML spec
	htmlElements := l.getHTMLElements()

	for _, element := range l.svgSpec.Elements {
		// Skip elements that already exist in HTML
		if htmlElements[element.Tag] {
			continue
		}
		// Use manual override for known void elements, otherwise use gostar data
		isVoid := element.NoChildren || svgVoidElements[element.Tag]

		spec := TagSpec{
			Name: element.Tag,
			Void: isVoid,
		}

		// Use a map to deduplicate attributes by key
		attrMap := make(map[string]Attribute)
		for _, attr := range element.Attributes {
			if attr.Key == "" {
				continue
			}

			// Skip global attributes - they'll be handled separately
			if _, isGlobal := globalAttrs[attr.Key]; isGlobal {
				continue
			}

			field := camelCase(attr.Key)
			attrType := "string"
			if l.isAttributeBoolean(attr) {
				attrType = "bool"
			}

			// Only add if not already seen (deduplication)
			if _, exists := attrMap[attr.Key]; !exists {
				attrMap[attr.Key] = Attribute{
					Field: field,
					Attr:  attr.Key,
					Type:  attrType,
				}
			}
		}

		// Convert map back to slice
		var elemAttributes []Attribute
		for _, attr := range attrMap {
			elemAttributes = append(elemAttributes, attr)
		}

		sort.Slice(elemAttributes, func(i, j int) bool {
			return elemAttributes[i].Attr < elemAttributes[j].Attr
		})

		spec.Attributes = elemAttributes
		specs = append(specs, spec)
	}

	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	return specs
}

// extractGlobalAttributes identifies attributes that appear in many SVG elements
// and should be treated as global attributes
func (l *Loader) extractGlobalAttributes() map[string]string {
	attrCounts := make(map[string]int)
	attrTypes := make(map[string]string)

	// Count how often each attribute appears
	for _, element := range l.svgSpec.Elements {
		for _, attr := range element.Attributes {
			if attr.Key != "" {
				attrCounts[attr.Key]++
				if attrTypes[attr.Key] == "" {
					if l.isAttributeBoolean(attr) {
						attrTypes[attr.Key] = "bool"
					} else {
						attrTypes[attr.Key] = "string"
					}
				}
			}
		}
	}

	// HTML global attributes that will be provided via embedded html.GlobalAttrs
	// We exclude these from SVG-specific globals to prevent duplicates
	htmlGlobals := map[string]bool{
		"class":           true,
		"id":              true,
		"style":           true,
		"accesskey":       true,
		"contenteditable": true,
		"dir":             true,
		"draggable":       true,
		"hidden":          true,
		"spellcheck":      true,
		"tabindex":        true,
		"title":           true,
	}

	// Consider attributes that appear in many elements as global
	totalElements := len(l.svgSpec.Elements)
	threshold := totalElements / 4 // 25% threshold
	if threshold < 3 {
		threshold = 3
	}

	globalAttrs := make(map[string]string)
	for attr, count := range attrCounts {
		// Skip HTML global attributes - they'll be handled by embedded html.GlobalAttrs
		if htmlGlobals[attr] {
			continue
		}
		if count >= threshold {
			globalAttrs[attr] = attrTypes[attr]
		}
	}

	// Include only SVG-specific global attributes (not covered by html.GlobalAttrs)
	svgSpecificGlobals := map[string]string{
		"fill":      "string",
		"stroke":    "string",
		"transform": "string",
	}

	for attr, typ := range svgSpecificGlobals {
		globalAttrs[attr] = typ
	}

	return globalAttrs
}

// getHTMLElements gets all HTML element tags from gostar to avoid conflicts
func (l *Loader) getHTMLElements() map[string]bool {
	htmlElements := make(map[string]bool)
	htmlSpec := cfg.HTML

	for _, element := range htmlSpec.Elements {
		if element.Tag != "" {
			htmlElements[element.Tag] = true
		}
	}

	return htmlElements
}

// camelCase converts kebab-case to CamelCase
func camelCase(name string) string {
	// Handle special cases for SVG attributes with dashes and mixed case
	delimiters := func(r rune) bool { return r == '-' || r == '_' }
	parts := strings.FieldsFunc(name, delimiters)

	if len(parts) == 0 {
		return name
	}

	// First part: capitalize first letter
	result := ""
	if len(parts[0]) > 0 {
		result = strings.ToUpper(parts[0][:1]) + parts[0][1:]
	}

	// Capitalize subsequent parts
	for i := 1; i < len(parts); i++ {
		p := parts[i]
		if len(p) > 0 {
			result += strings.ToUpper(p[:1]) + p[1:]
		}
	}

	return result
}
