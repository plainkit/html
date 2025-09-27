package spec

import (
	"fmt"
	"sort"
	"strings"

	"github.com/delaneyj/gostar/cfg"
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

// Loader handles loading and parsing of HTML specification from gostar
type Loader struct {
	htmlSpec *pb.Namespace
}

// NewLoader creates a new spec loader using gostar data
func NewLoader(specsDir string) *Loader {
	return &Loader{htmlSpec: cfg.HTML}
}

// getGostarSpec returns the gostar HTML specification
func (l *Loader) getGostarSpec() *pb.Namespace {
	return l.htmlSpec
}

// isVoidElement checks if an element is self-closing based on gostar data
func (l *Loader) isVoidElement(tagName string) bool {
	for _, element := range l.htmlSpec.Elements {
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

// LoadAllTagSpecsFromGostar loads all tag specifications from gostar data
func (l *Loader) LoadAllTagSpecsFromGostar() ([]TagSpec, error) {
	return l.convertGostarToTagSpecs(), nil
}

// LoadGlobalAttributesFromGostar loads global attributes from gostar data
func (l *Loader) LoadGlobalAttributesFromGostar() ([]Attribute, error) {
	// Note: gostar doesn't seem to have explicit global attributes
	// We'll collect common attributes that appear across many elements
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

func (l *Loader) convertGostarToTagSpecs() []TagSpec {
	var specs []TagSpec

	globalAttrs := l.extractGlobalAttributes()

	for _, element := range l.htmlSpec.Elements {
		// Skip SVG tag - it's handled by the separate SVG package
		if element.Tag == "svg" {
			continue
		}

		spec := TagSpec{
			Name:          element.Tag,
			Void:          element.NoChildren,
			ParentTargets: []string{}, // gostar doesn't provide parent-child constraints
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

// extractGlobalAttributes identifies attributes that appear in many elements
// and should be treated as global attributes
func (l *Loader) extractGlobalAttributes() map[string]string {
	attrCounts := make(map[string]int)
	attrTypes := make(map[string]string)
	totalElements := len(l.htmlSpec.Elements)

	// Count how often each attribute appears
	for _, element := range l.htmlSpec.Elements {
		for _, attr := range element.Attributes {
			attrCounts[attr.Key]++
			if l.isAttributeBoolean(attr) {
				attrTypes[attr.Key] = "bool"
			} else {
				attrTypes[attr.Key] = "string"
			}
		}
	}

	// Consider attributes that appear in at least 10% of elements as global
	threshold := totalElements / 10
	globalAttrs := make(map[string]string)
	for attrName, count := range attrCounts {
		if count >= threshold {
			globalAttrs[attrName] = attrTypes[attrName]
		}
	}

	return globalAttrs
}

// CollectAllAttributes collects unique attributes from all tag specs
func (l *Loader) CollectAllAttributes(specs []TagSpec) map[string]Attribute {
	allAttributes := make(map[string]Attribute)

	for _, spec := range specs {
		for _, attr := range spec.Attributes {
			key := strings.ToLower(attr.Attr)
			if existing, exists := allAttributes[key]; exists {
				if existing.Type == "bool" && attr.Type == "string" {
					allAttributes[key] = attr
				}
			} else {
				allAttributes[key] = attr
			}
		}
	}

	return allAttributes
}

// CamelCase converts kebab-case to CamelCase
func CamelCase(name string) string {
	return camelCase(name)
}

func camelCase(name string) string {
	cleaned := strings.ReplaceAll(strings.ReplaceAll(name, "-", ""), "_", "")
	if len(cleaned) == 0 {
		return ""
	}
	return strings.ToUpper(cleaned[:1]) + cleaned[1:]
}
