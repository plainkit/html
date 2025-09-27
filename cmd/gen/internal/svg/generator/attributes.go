package generator

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/delaneyj/gostar/cfg"
	"github.com/plainkit/html/cmd/gen/internal/svg/spec"
	"github.com/plainkit/html/cmd/gen/internal/svg/utils"
)

// AttributesGenerator generates centralized SVG attributes file
type AttributesGenerator struct{}

// NewAttributesGenerator creates a new attributes generator
func NewAttributesGenerator() *AttributesGenerator {
	return &AttributesGenerator{}
}

// AttributesTemplateData holds data for template rendering
type AttributesTemplateData struct {
	Attributes []AttributeData
}

// AttributeData represents a single attribute for template rendering
type AttributeData struct {
	Field  string // Go field name (e.g., "D", "Fill")
	Type   string // Attribute type ("string" or "bool")
	Attr   string // HTML attribute name (e.g., "d", "fill")
	GoType string // Go type ("string" or "bool")
}

// GenerateSource creates the source code for centralized SVG attributes file
func (g *AttributesGenerator) GenerateSource(attributes map[string]spec.Attribute) string {
	// Get HTML attributes to exclude them from SVG generation
	htmlAttrs := g.getHTMLAttributes()

	var keys []string
	for key := range attributes {
		// Skip HTML attributes to avoid conflicts
		if _, isHTMLAttr := htmlAttrs[key]; isHTMLAttr {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var templateAttrs []AttributeData
	for _, key := range keys {
		attr := attributes[key]

		templateAttrs = append(templateAttrs, AttributeData{
			Field:  attr.Field,
			Type:   attr.Type,
			Attr:   attr.Attr,
			GoType: utils.GoType(attr.Type),
		})
	}

	data := AttributesTemplateData{
		Attributes: templateAttrs,
	}

	tmpl, err := template.New("svgAttributes").Parse(attributesTemplate)
	if err != nil {
		panic("failed to parse SVG attributes template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute SVG attributes template: " + err.Error())
	}

	return buf.String()
}

// getHTMLAttributes extracts HTML attributes from gostar to exclude them from SVG generation
func (g *AttributesGenerator) getHTMLAttributes() map[string]bool {
	htmlSpec := cfg.HTML
	attrCounts := make(map[string]int)
	totalElements := len(htmlSpec.Elements)

	// Count how often each attribute appears in HTML
	for _, element := range htmlSpec.Elements {
		for _, attr := range element.Attributes {
			if attr.Key != "" {
				attrCounts[attr.Key]++
			}
		}
	}

	// Consider attributes that appear in 90%+ of elements, plus known HTML global attributes
	threshold := totalElements * 9 / 10
	htmlAttrs := make(map[string]bool)

	// Add attributes that appear frequently in HTML
	for attrName, count := range attrCounts {
		if count >= threshold {
			htmlAttrs[attrName] = true
		}
	}

	// Always include known HTML global attributes (from HTML spec) to avoid conflicts
	knownHTMLGlobals := []string{
		"class", "id", "style", "title", "contenteditable", "draggable",
		"hidden", "lang", "spellcheck", "tabindex", "accesskey", "dir",
	}

	for _, attr := range knownHTMLGlobals {
		htmlAttrs[attr] = true
	}

	// Also add all HTML attributes that could conflict with SVG ones
	// This includes any attribute that appears in HTML regardless of frequency
	for attrName := range attrCounts {
		htmlAttrs[attrName] = true
	}

	return htmlAttrs
}
