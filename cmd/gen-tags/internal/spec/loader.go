package spec

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// VoidElements defines which HTML elements are self-closing
var VoidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true,
	"embed": true, "hr": true, "img": true, "input": true,
	"link": true, "meta": true, "param": true, "source": true,
	"track": true, "wbr": true,
}

// BoolAttributes defines which HTML attributes are boolean
var BoolAttributes = map[string]bool{
	"allowfullscreen": true,
	"async":           true,
	"autofocus":       true,
	"autoplay":        true,
	"checked":         true,
	"contenteditable": true,
	"controls":        true,
	"default":         true,
	"defer":           true,
	"disabled":        true,
	"formnovalidate":  true,
	"hidden":          true,
	"loop":            true,
	"multiple":        true,
	"muted":           true,
	"open":            true,
	"readonly":        true,
	"required":        true,
	"reversed":        true,
	"selected":        true,
}

// ParentMap defines parent-child relationships
var ParentMap = map[string][]string{}

// Loader handles loading and parsing of HTML specification files
type Loader struct {
	specsDir string
}

// WooormAttributeData represents the structure from wooorm/html-element-attributes
type WooormAttributeData map[string][]string

// StandardHTML5Elements contains all standard HTML5 elements that should be included
var StandardHTML5Elements = map[string]bool{
	// Document metadata
	"title": true,

	// Sectioning root
	"html": true,
	"body": true,

	// Content sectioning
	"article": true,
	"aside":   true,
	"footer":  true,
	"header":  true,
	"main":    true,
	"nav":     true,
	"section": true,
	"search":  true,

	// Text content
	"blockquote": true,
	"dd":         true,
	"div":        true,
	"dl":         true,
	"dt":         true,
	"figcaption": true,
	"figure":     true,
	"hr":         true,
	"li":         true,
	"menu":       true,
	"ol":         true,
	"p":          true,
	"pre":        true,
	"ul":         true,

	// Inline text semantics
	"a":      true,
	"abbr":   true,
	"b":      true,
	"bdi":    true,
	"bdo":    true,
	"br":     true,
	"cite":   true,
	"code":   true,
	"data":   true,
	"dfn":    true,
	"em":     true,
	"i":      true,
	"kbd":    true,
	"mark":   true,
	"q":      true,
	"rp":     true,
	"rt":     true,
	"ruby":   true,
	"s":      true,
	"samp":   true,
	"small":  true,
	"span":   true,
	"strong": true,
	"sub":    true,
	"sup":    true,
	"time":   true,
	"u":      true,
	"var":    true,
	"wbr":    true,

	// Image and multimedia
	"area":    true,
	"audio":   true,
	"img":     true,
	"map":     true,
	"picture": true,
	"source":  true,
	"track":   true,
	"video":   true,

	// Embedded content
	"canvas": true,
	"embed":  true,
	"iframe": true,
	"object": true,
	"slot":   true,
	"svg":    true,

	// Forms
	"button":   true,
	"datalist": true,
	"fieldset": true,
	"form":     true,
	"input":    true,
	"label":    true,
	"legend":   true,
	"meter":    true,
	"optgroup": true,
	"option":   true,
	"output":   true,
	"progress": true,
	"select":   true,
	"textarea": true,

	// Interactive elements
	"details": true,
	"dialog":  true,
	"summary": true,

	// Web Components
	"template": true,

	// Scripting
	"noscript": true,
	"script":   true,

	// Document metadata
	"base":  true,
	"head":  true,
	"link":  true,
	"meta":  true,
	"style": true,

	// Table content
	"caption":  true,
	"col":      true,
	"colgroup": true,
	"table":    true,
	"tbody":    true,
	"td":       true,
	"tfoot":    true,
	"th":       true,
	"thead":    true,
	"tr":       true,

	// Demarcating edits
	"del": true,
	"ins": true,
}

// NewLoader creates a new spec loader for the given directory
func NewLoader(specsDir string) *Loader {
	return &Loader{specsDir: specsDir}
}

// FetchWooormData fetches HTML element attributes from the wooorm GitHub repository
func (l *Loader) FetchWooormData() (WooormAttributeData, error) {
	url := "https://raw.githubusercontent.com/wooorm/html-element-attributes/refs/heads/main/index.js"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch wooorm data: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch wooorm data: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read wooorm data: %w", err)
	}

	return parseWooormData(string(body))
}

// parseWooormData parses the JavaScript file content to extract HTML element attributes
func parseWooormData(content string) (WooormAttributeData, error) {
	startPattern := `export const htmlElementAttributes = {`
	startIndex := strings.Index(content, startPattern)
	if startIndex == -1 {
		return nil, fmt.Errorf("could not find 'export const htmlElementAttributes = {' in content")
	}

	braceStart := startIndex + len(startPattern) - 1
	braceCount := 0
	endIndex := -1

LoopEnd:
	for i := braceStart; i < len(content); i++ {
		switch content[i] {
		case '{':
			braceCount++
		case '}':
			braceCount--
			if braceCount == 0 {
				endIndex = i + 1
				break LoopEnd
			}
		}
	}

	if endIndex == -1 {
		return nil, fmt.Errorf("could not find matching closing brace for htmlElementAttributes object")
	}

	objectStr := content[braceStart:endIndex]

	jsonStr := convertJSObjectToJSON(objectStr)

	var data WooormAttributeData
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, fmt.Errorf("parse wooorm data as JSON: %w", err)
	}

	return data, nil
}

// convertJSObjectToJSON converts JavaScript object notation to valid JSON
func convertJSObjectToJSON(jsObject string) string {
	result := strings.ReplaceAll(jsObject, "'", "\"")

	keyRegex := regexp.MustCompile(`([{,]\s*)([a-zA-Z_$][a-zA-Z0-9_$]*)\s*:`)
	result = keyRegex.ReplaceAllString(result, `${1}"${2}":`)

	return result
}

func (l *Loader) LoadAllTagSpecsFromWooorm() ([]TagSpec, error) {
	wooormData, err := l.FetchWooormData()
	if err != nil {
		return nil, fmt.Errorf("fetch wooorm data: %w", err)
	}

	return l.convertWooormToTagSpecs(wooormData), nil
}

func (l *Loader) LoadGlobalAttributesFromWooorm() ([]Attribute, error) {
	wooormData, err := l.FetchWooormData()
	if err != nil {
		return nil, fmt.Errorf("fetch wooorm data: %w", err)
	}

	globalAttrs, exists := wooormData["*"]
	if !exists {
		return nil, fmt.Errorf("no global attributes found in wooorm data")
	}

	var attributes []Attribute
	for _, attrName := range globalAttrs {
		if attrName == "" {
			continue
		}

		field := camelCase(attrName)
		attr := Attribute{
			Field: field,
			Attr:  attrName,
			Type:  "string", // default type
		}

		// Check if it's a boolean attribute
		if BoolAttributes[attrName] {
			attr.Type = "bool"
		}

		attributes = append(attributes, attr)
	}

	// Sort attributes by name for deterministic output
	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Attr < attributes[j].Attr
	})

	return attributes, nil
}

// convertWooormToTagSpecs converts wooorm data to TagSpec format and adds missing standard HTML5 elements
func (l *Loader) convertWooormToTagSpecs(wooormData WooormAttributeData) []TagSpec {
	var specs []TagSpec
	processedElements := make(map[string]bool)

	// Get global attributes to exclude from element-specific attributes
	globalAttrs := make(map[string]bool)
	if globals, exists := wooormData["*"]; exists {
		for _, attr := range globals {
			globalAttrs[attr] = true
		}
	}

	// Convert each element from wooorm data
	for tagName, attributes := range wooormData {
		if tagName == "*" {
			continue // Skip global attributes
		}

		spec := TagSpec{
			Name:          tagName,
			Void:          VoidElements[tagName],
			ParentTargets: ParentMap[tagName],
		}

		// Convert attributes, excluding global ones
		var elemAttributes []Attribute
		for _, attrName := range attributes {
			if attrName == "" || globalAttrs[attrName] {
				continue // Skip empty or global attributes
			}

			field := camelCase(attrName)
			attr := Attribute{
				Field: field,
				Attr:  attrName,
				Type:  "string", // default type
			}

			if BoolAttributes[strings.ToLower(attrName)] {
				attr.Type = "bool"
			}

			elemAttributes = append(elemAttributes, attr)
		}

		// Sort attributes by name for deterministic output
		sort.Slice(elemAttributes, func(i, j int) bool {
			return elemAttributes[i].Attr < elemAttributes[j].Attr
		})

		spec.Attributes = elemAttributes
		specs = append(specs, spec)
		processedElements[tagName] = true
	}

	// Add missing standard HTML5 elements that weren't in wooorm data
	for elementName := range StandardHTML5Elements {
		if !processedElements[elementName] {
			// Create a spec for the missing element with no element-specific attributes
			spec := TagSpec{
				Name:          elementName,
				Void:          VoidElements[elementName],
				ParentTargets: ParentMap[elementName],
				Attributes:    []Attribute{}, // No element-specific attributes
			}
			specs = append(specs, spec)
		}
	}

	// Sort specs by tag name for deterministic output
	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	return specs
}

// LoadGlobalAttributes loads the global attributes specification
func (l *Loader) LoadGlobalAttributes() (*GlobalAttributesSpec, error) {
	globalPath := filepath.Join(l.specsDir, "global_attributes.json")
	data, err := os.ReadFile(globalPath)
	if err != nil {
		return nil, fmt.Errorf("read global attributes spec: %w", err)
	}

	var spec GlobalAttributesSpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("unmarshal global attributes: %w", err)
	}

	return &spec, nil
}

// LoadTagSpec loads specification for a single HTML tag
func (l *Loader) LoadTagSpec(filename, tagName string) (TagSpec, error) {
	var spec TagSpec
	spec.Name = tagName
	spec.Void = VoidElements[tagName]
	spec.ParentTargets = ParentMap[tagName]

	path := filepath.Join(l.specsDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return spec, fmt.Errorf("read spec file: %w", err)
	}

	var doc BrowserCompatData
	if err := json.Unmarshal(data, &doc); err != nil {
		return spec, fmt.Errorf("unmarshal spec: %w", err)
	}

	elemData, ok := doc.HTML.Elements[tagName]
	if !ok {
		return spec, nil
	}

	// Extract and sort attribute names for deterministic output
	var keys []string
	for k := range elemData {
		if k == "__compat" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create attribute specs
	for _, key := range keys {
		field := camelCase(key)
		attr := Attribute{
			Field: field,
			Attr:  key,
			Type:  "string", // default
		}

		if BoolAttributes[strings.ToLower(key)] {
			attr.Type = "bool"
		}

		spec.Attributes = append(spec.Attributes, attr)
	}

	return spec, nil
}

// LoadAllTagSpecs loads all tag specifications from the specs directory
func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	entries, err := os.ReadDir(l.specsDir)
	if err != nil {
		return nil, fmt.Errorf("read specs directory: %w", err)
	}

	var specs []TagSpec
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".json") || name == "global_attributes.json" {
			continue
		}

		tagName := strings.TrimSuffix(name, ".json")
		spec, err := l.LoadTagSpec(name, tagName)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}

		specs = append(specs, spec)
	}

	return specs, nil
}

// CollectAllAttributes collects unique attributes from all tag specs
func (l *Loader) CollectAllAttributes(specs []TagSpec) map[string]Attribute {
	allAttributes := make(map[string]Attribute)

	for _, spec := range specs {
		for _, attr := range spec.Attributes {
			key := strings.ToLower(attr.Attr)
			if existing, exists := allAttributes[key]; exists {
				// Keep string type if there's a conflict (more permissive)
				if existing.Type == "bool" && attr.Type == "string" {
					allAttributes[key] = attr
				}
			} else {
				allAttributes[key] = attr
			}
		}
	}

	return allAttributes
}

// CamelCase converts kebab-case to CamelCase (exported for use by main)
func CamelCase(name string) string {
	return camelCase(name)
}

// camelCase converts kebab-case to simple CamelCase (single word approach)
func camelCase(name string) string {
	// Remove delimiters and join into single word, then capitalize first letter
	cleaned := strings.ReplaceAll(strings.ReplaceAll(name, "-", ""), "_", "")
	if len(cleaned) == 0 {
		return ""
	}
	return strings.ToUpper(cleaned[:1]) + cleaned[1:]
}
