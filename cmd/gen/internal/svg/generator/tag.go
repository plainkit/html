package generator

import (
	"bytes"
	"text/template"

	"github.com/plainkit/html/cmd/gen/internal/svg/spec"
	"github.com/plainkit/html/cmd/gen/internal/svg/utils"
)

// TagGenerator generates individual SVG tag files
type TagGenerator struct{}

// NewTagGenerator creates a new tag generator
func NewTagGenerator() *TagGenerator {
	return &TagGenerator{}
}

// TagTemplateData holds data for tag template rendering
type TagTemplateData struct {
	Name         string             // SVG element name (e.g., "rect", "circle")
	Title        string             // CamelCase version for function name (e.g., "Rect", "Circle")
	StructName   string             // Struct name (e.g., "RectAttrs")
	ArgInterface string             // Arg interface name (e.g., "RectArg")
	Void         bool               // Whether the element is self-closing
	Attributes   []TagAttributeData // Element-specific attributes
}

// TagAttributeData represents attribute data for tag templates
type TagAttributeData struct {
	Field  string // Go field name (e.g., "Width")
	Type   string // Attribute type ("string" or "bool")
	Attr   string // HTML attribute name (e.g., "width")
	GoType string // Go type ("string" or "bool")
}

// GenerateSource generates the Go source code for a single SVG tag
func (g *TagGenerator) GenerateSource(tagSpec spec.TagSpec) string {
	title := utils.CamelCase(tagSpec.Name)
	// Special case: for the "svg" element, use empty title to avoid SvgSvg duplication
	if tagSpec.Name == "svg" {
		title = ""
	}

	structName := title + "Attrs"

	// Convert attributes to template data
	var templateAttrs []TagAttributeData
	for _, attr := range tagSpec.Attributes {
		templateAttrs = append(templateAttrs, TagAttributeData{
			Field:  attr.Field,
			Type:   attr.Type,
			Attr:   attr.Attr,
			GoType: utils.GoType(attr.Type),
		})
	}

	// Prepare template data
	argInterface := title + "Arg"
	// Special case: for the "svg" element, use "Arg" to avoid empty + "Arg" = "Arg"
	if tagSpec.Name == "svg" {
		argInterface = "Arg"
	}

	data := TagTemplateData{
		Name:         tagSpec.Name,
		Title:        title,
		StructName:   structName,
		ArgInterface: argInterface,
		Void:         tagSpec.Void,
		Attributes:   templateAttrs,
	}

	// Parse and execute template
	tmpl, err := template.New("tag").Parse(tagTemplate)
	if err != nil {
		panic("failed to parse tag template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute tag template: " + err.Error())
	}

	return buf.String()
}
