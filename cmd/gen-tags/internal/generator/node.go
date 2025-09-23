package generator

import (
	"bytes"
	"text/template"

	"github.com/plainkit/html/cmd/gen-tags/internal/spec"
	"github.com/plainkit/html/cmd/gen-tags/internal/utils"
)

// NodeGenerator handles generation and updating of Node apply methods
type NodeGenerator struct{}

// NewNodeGenerator creates a new node generator
func NewNodeGenerator() *NodeGenerator {
	return &NodeGenerator{}
}

// NodeTemplateData holds data for node template
type NodeTemplateData struct {
	TagNames []string
}

// GenerateCompleteFile creates the entire core_node.go file from scratch
func (g *NodeGenerator) GenerateCompleteFile(specs []spec.TagSpec) string {
	// Collect all unique tag names and deduplicate
	tagNames := make(map[string]bool)
	for _, spec := range specs {
		tagNames[spec.Name] = true
	}

	// Convert to sorted slice for consistent output
	var sortedTags []string
	for tagName := range tagNames {
		sortedTags = append(sortedTags, tagName)
	}

	// Sort for consistent output
	for i := 0; i < len(sortedTags)-1; i++ {
		for j := i + 1; j < len(sortedTags); j++ {
			if sortedTags[i] > sortedTags[j] {
				sortedTags[i], sortedTags[j] = sortedTags[j], sortedTags[i]
			}
		}
	}

	// Convert tag names to CamelCase
	var camelCaseNames []string
	for _, tagName := range sortedTags {
		camelCaseNames = append(camelCaseNames, utils.CamelCase(tagName))
	}

	data := NodeTemplateData{
		TagNames: camelCaseNames,
	}

	// Execute template
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
