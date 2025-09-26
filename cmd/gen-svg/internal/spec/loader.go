package spec

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

// VoidElements defines SVG elements that are self-closing
var VoidElements = map[string]bool{
	"animateMotion":    true,
	"animateTransform": true,
	"circle":           true,
	"ellipse":          true,
	"feDistantLight":   true,
	"feDropShadow":     true,
	"feFuncA":          true,
	"feFuncB":          true,
	"feFuncG":          true,
	"feFuncR":          true,
	"fePointLight":     true,
	"feSpotLight":      true,
	"line":             true,
	"path":             true,
	"polygon":          true,
	"polyline":         true,
	"rect":             true,
	"stop":             true,
	"use":              true,
}

// BoolAttributes defines SVG attributes that are boolean
var BoolAttributes = map[string]bool{
	"externalResourcesRequired": true,
	"focusable":                 true,
	"crossorigin":               true,
}

// Loader handles loading and parsing of SVG specification
type Loader struct{}

// NewLoader creates a new spec loader
func NewLoader() *Loader {
	return &Loader{}
}

// FetchSvgElementAttributes fetches SVG element attributes from the wooorm GitHub repository
func (l *Loader) FetchSvgElementAttributes() (SvgElementAttributes, error) {
	url := "https://raw.githubusercontent.com/wooorm/svg-element-attributes/refs/heads/main/index.js"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch svg-element-attributes: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch svg-element-attributes: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read svg-element-attributes: %w", err)
	}

	return l.parseSvgElementAttributes(string(body))
}

// parseSvgElementAttributes parses the JavaScript module content
func (l *Loader) parseSvgElementAttributes(content string) (SvgElementAttributes, error) {
	// Find the export statement with the object
	re := regexp.MustCompile(`export const svgElementAttributes = ({[\s\S]*?})\s*;?\s*$`)
	matches := re.FindStringSubmatch(content)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not find svgElementAttributes export in content")
	}

	jsonStr := matches[1]

	// Convert JavaScript object to JSON
	jsonStr = l.convertJSObjectToJSON(jsonStr)

	var data SvgElementAttributes
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, fmt.Errorf("parse svg-element-attributes as JSON: %w", err)
	}

	return data, nil
}

// convertJSObjectToJSON converts a JavaScript object literal to valid JSON
func (l *Loader) convertJSObjectToJSON(jsObject string) string {
	// Replace single quotes with double quotes
	result := strings.ReplaceAll(jsObject, "'", "\"")

	// Add quotes around unquoted keys
	keyRegex := regexp.MustCompile(`(?m)^\s*([a-zA-Z_$][a-zA-Z0-9_$-]*)\s*:`)
	result = keyRegex.ReplaceAllString(result, "\"$1\":")

	return result
}

// LoadAllTagSpecs loads all SVG tag specifications from the fetched data
func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	fmt.Println("Fetching SVG element attributes from wooorm repository...")
	svgData, err := l.FetchSvgElementAttributes()
	if err != nil {
		return nil, fmt.Errorf("fetch svg data: %w", err)
	}
	fmt.Printf("Loaded %d SVG element specifications\n", len(svgData))

	return l.convertToTagSpecs(svgData), nil
}

// convertToTagSpecs converts SvgElementAttributes to TagSpec slice
func (l *Loader) convertToTagSpecs(svgData SvgElementAttributes) []TagSpec {
	var specs []TagSpec

	// Get global attributes
	globalAttrs := make(map[string]bool)
	if globalAttrList, exists := svgData["*"]; exists {
		for _, attr := range globalAttrList {
			globalAttrs[attr] = true
		}
	}

	// Process each element
	for tagName, attributes := range svgData {
		if tagName == "*" {
			continue // Skip global attributes entry
		}

		spec := TagSpec{
			Name: tagName,
			Void: VoidElements[tagName],
		}

		// Add element-specific attributes (excluding globals)
		for _, attrName := range attributes {
			if attrName == "" || globalAttrs[attrName] {
				continue
			}

			field := CamelCase(attrName)
			attr := Attribute{
				Field: field,
				Attr:  attrName,
				Type:  "string",
			}

			if BoolAttributes[strings.ToLower(attrName)] {
				attr.Type = "bool"
			}

			spec.Attributes = append(spec.Attributes, attr)
		}

		// Sort attributes by field name for consistency
		sort.Slice(spec.Attributes, func(i, j int) bool {
			return spec.Attributes[i].Field < spec.Attributes[j].Field
		})

		specs = append(specs, spec)
	}

	// Sort specs by element name
	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	return specs
}

// CamelCase converts kebab-case to CamelCase
func CamelCase(name string) string {
	// Handle special cases for SVG attributes with dashes and mixed case
	delimiters := func(r rune) bool { return r == '-' || r == '_' }
	parts := strings.FieldsFunc(name, delimiters)

	if len(parts) == 0 {
		return name
	}

	// First part: capitalize first letter
	result := ""
	if len(parts[0]) > 0 {
		result = strings.ToUpper(parts[0][:1]) + parts[0][1:]
	}

	// Capitalize subsequent parts
	for i := 1; i < len(parts); i++ {
		p := parts[i]
		if len(p) > 0 {
			result += strings.ToUpper(p[:1]) + p[1:]
		}
	}

	return result
}
