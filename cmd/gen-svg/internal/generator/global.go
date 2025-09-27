package generator

import (
	"bytes"
	"text/template"

	"github.com/plainkit/html/cmd/gen-svg/internal/spec"
	"github.com/plainkit/html/cmd/gen-svg/internal/utils"
)

// GlobalGenerator generates SVG global attributes file
type GlobalGenerator struct{}

// NewGlobalGenerator creates a new global generator
func NewGlobalGenerator() *GlobalGenerator {
	return &GlobalGenerator{}
}

// GlobalTemplateData holds data for global template rendering
type GlobalTemplateData struct {
	Attributes []GlobalAttributeData
}

// GlobalAttributeData represents a global attribute for template rendering
type GlobalAttributeData struct {
	Field  string // Go field name (e.g., "Class")
	Type   string // Attribute type ("string" or "bool")
	Attr   string // HTML attribute name (e.g., "class")
	GoType string // Go type ("string" or "bool")
}

// GenerateSource generates the Go source code for SVG global attributes
func (g *GlobalGenerator) GenerateSource(attributes []spec.Attribute) string {
	var templateAttrs []GlobalAttributeData
	for _, attr := range attributes {
		templateAttrs = append(templateAttrs, GlobalAttributeData{
			Field:  attr.Field,
			Type:   attr.Type,
			Attr:   attr.Attr,
			GoType: utils.GoType(attr.Type),
		})
	}

	data := GlobalTemplateData{
		Attributes: templateAttrs,
	}

	tmpl, err := template.New("svgGlobal").Parse(globalTemplate)
	if err != nil {
		panic("failed to parse SVG global template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute SVG global template: " + err.Error())
	}

	return buf.String()
}
