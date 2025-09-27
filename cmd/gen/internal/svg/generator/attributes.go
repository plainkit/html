package generator

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"github.com/plainkit/html/cmd/gen/internal/svg/spec"
	"github.com/plainkit/html/cmd/gen/internal/svg/utils"
	"github.com/plainkit/tags"
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
	htmlAttrs := make(map[string]bool)

	for _, attr := range tags.HTML.Globals {
		name := strings.ToLower(attr.Name)
		if name != "" {
			htmlAttrs[name] = true
		}
	}

	for _, element := range tags.HTML.Elements {
		for _, attr := range element.Attributes {
			name := strings.ToLower(attr.Name)
			if name != "" {
				htmlAttrs[name] = true
			}
		}
	}

	return htmlAttrs
}
