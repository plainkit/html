package generator

import (
	"bytes"
	"text/template"

	"github.com/plainkit/html/cmd/gen-tags/internal/spec"
	"github.com/plainkit/html/cmd/gen-tags/internal/utils"
)

// TagGenerator generates individual tag files
type TagGenerator struct{}

// NewTagGenerator creates a new tag generator
func NewTagGenerator() *TagGenerator {
	return &TagGenerator{}
}

// TagTemplateData holds data for tag template rendering
type TagTemplateData struct {
	Name         string
	Title        string
	StructName   string
	ArgInterface string
	Void         bool
	Attributes   []TagAttributeData
}

// TagAttributeData represents attribute data for tag templates
type TagAttributeData struct {
	Field  string
	Type   string
	Attr   string
	GoType string
}

func (g *TagGenerator) GenerateSource(tagSpec spec.TagSpec) string {
	title := utils.CamelCase(tagSpec.Name)
	structName := title + "Attrs"
	argInterface := title + "Arg"

	var templateAttrs []TagAttributeData
	for _, attr := range tagSpec.Attributes {
		templateAttrs = append(templateAttrs, TagAttributeData{
			Field:  attr.Field,
			Type:   attr.Type,
			Attr:   attr.Attr,
			GoType: utils.GoType(attr.Type),
		})
	}

	data := TagTemplateData{
		Name:         tagSpec.Name,
		Title:        title,
		StructName:   structName,
		ArgInterface: argInterface,
		Void:         tagSpec.Void,
		Attributes:   templateAttrs,
	}

	tmpl, err := template.New("tag").Parse(tagsTemplate)
	if err != nil {
		panic("failed to parse tags template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute tags template: " + err.Error())
	}

	return buf.String()
}
