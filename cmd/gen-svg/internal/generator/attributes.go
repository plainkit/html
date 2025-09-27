package generator

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/plainkit/html/cmd/gen-svg/internal/spec"
	"github.com/plainkit/html/cmd/gen-svg/internal/utils"
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
	var keys []string
	for key := range attributes {
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
