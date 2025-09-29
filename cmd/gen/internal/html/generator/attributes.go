package generator

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/plainkit/html/cmd/gen/internal/html/spec"
	"github.com/plainkit/html/cmd/gen/internal/html/utils"
)

// AttributesGenerator generates centralized attributes file
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
	Field  string
	Type   string
	Attr   string
	GoType string
}

// GenerateSource creates the source code for centralized attributes file
func (g *AttributesGenerator) GenerateSource(attributes map[string]spec.Attribute) string {
	var keys []string
	for key := range attributes {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var templateAttrs []AttributeData

	for _, key := range keys {
		attr := attributes[key]

		if attr.Attr == "data" {
			continue // AData is handled in core_global.go as AData(k, v string)
		}

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

	tmpl, err := template.New("attributes").Parse(attributesTemplate)
	if err != nil {
		panic("failed to parse attributes template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute attributes template: " + err.Error())
	}

	return buf.String()
}
