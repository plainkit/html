package generator

import (
	"bytes"
	"text/template"

	"github.com/plainkit/html/cmd/gen/internal/html/spec"
	"github.com/plainkit/html/cmd/gen/internal/html/utils"
)

// NodeGenerator generates core_node.go file
type NodeGenerator struct{}

// NewNodeGenerator creates a new node generator
func NewNodeGenerator() *NodeGenerator {
	return &NodeGenerator{}
}

// NodeTemplateData holds data for node template rendering
type NodeTemplateData struct {
	TagNames []string
}

// GenerateCompleteFile creates complete core_node.go file content
func (g *NodeGenerator) GenerateCompleteFile(specs []spec.TagSpec) string {
	tagNames := make(map[string]bool)
	for _, spec := range specs {
		tagNames[spec.Name] = true
	}

	var sortedTags []string
	for tagName := range tagNames {
		sortedTags = append(sortedTags, tagName)
	}

	for i := 0; i < len(sortedTags)-1; i++ {
		for j := i + 1; j < len(sortedTags); j++ {
			if sortedTags[i] > sortedTags[j] {
				sortedTags[i], sortedTags[j] = sortedTags[j], sortedTags[i]
			}
		}
	}

	var camelCaseNames []string
	for _, tagName := range sortedTags {
		camelCaseNames = append(camelCaseNames, utils.CamelCase(tagName))
	}

	data := NodeTemplateData{
		TagNames: camelCaseNames,
	}

	tmpl, err := template.New("node").Parse(nodeTemplate)
	if err != nil {
		panic("failed to parse node template: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic("failed to execute node template: " + err.Error())
	}

	return buf.String()
}
