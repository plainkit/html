package generator

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/plainkit/html/cmd/gen-tags/internal/spec"
	"github.com/plainkit/html/cmd/gen-tags/internal/utils"
)

// AttributesGenerator handles generation of the centralized attributes file
type AttributesGenerator struct{}

// NewAttributesGenerator creates a new attributes generator
func NewAttributesGenerator() *AttributesGenerator {
	return &AttributesGenerator{}
}

// TemplateData holds the data for the attributes template
type AttributesTemplateData struct {
	Attributes []AttributeData
}

type AttributeData struct {
	Field  string
	Type   string
	Attr   string
	GoType string
}

// GenerateSource creates the source code for the attributes file
func (g *AttributesGenerator) GenerateSource(attributes map[string]spec.Attribute) string {
	// Sort attributes for deterministic output
	var keys []string
	for key := range attributes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Prepare template data
	var templateAttrs []AttributeData
	for _, key := range keys {
		attr := attributes[key]

		// Skip attributes that are handled in core_global.go
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

	// Execute template
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
