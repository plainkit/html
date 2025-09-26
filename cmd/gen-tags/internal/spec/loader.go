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

// VoidElements defines HTML elements that are self-closing
var VoidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true,
	"embed": true, "hr": true, "img": true, "input": true,
	"link": true, "meta": true, "param": true, "source": true,
	"track": true, "wbr": true,
}

// BoolAttributes defines HTML attributes that are boolean
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

// StandardHTML5Elements contains all standard HTML5 elements
var StandardHTML5Elements = map[string]bool{
	"title": true,

	"html": true,
	"body": true,

	"article": true,
	"aside":   true,
	"footer":  true,
	"header":  true,
	"main":    true,
	"nav":     true,
	"section": true,
	"search":  true,

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

	"area":    true,
	"audio":   true,
	"img":     true,
	"map":     true,
	"picture": true,
	"source":  true,
	"track":   true,
	"video":   true,

	"canvas": true,
	"embed":  true,
	"iframe": true,
	"object": true,
	"slot":   true,
	"svg":    true,

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

	"details": true,
	"dialog":  true,
	"summary": true,

	"template": true,

	"noscript": true,
	"script":   true,

	"base":  true,
	"head":  true,
	"link":  true,
	"meta":  true,
	"style": true,

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

func convertJSObjectToJSON(jsObject string) string {
	result := strings.ReplaceAll(jsObject, "'", "\"")

	keyRegex := regexp.MustCompile(`([{,]\s*)([a-zA-Z_$][a-zA-Z0-9_$]*)\s*:`)
	result = keyRegex.ReplaceAllString(result, `${1}"${2}":`)

	return result
}

// LoadAllTagSpecsFromWooorm loads all tag specifications from wooorm data
func (l *Loader) LoadAllTagSpecsFromWooorm() ([]TagSpec, error) {
	wooormData, err := l.FetchWooormData()
	if err != nil {
		return nil, fmt.Errorf("fetch wooorm data: %w", err)
	}

	return l.convertWooormToTagSpecs(wooormData), nil
}

// LoadGlobalAttributesFromWooorm loads global attributes from wooorm data
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
			Type:  "string",
		}

		if BoolAttributes[attrName] {
			attr.Type = "bool"
		}

		attributes = append(attributes, attr)
	}

	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Attr < attributes[j].Attr
	})

	return attributes, nil
}

func (l *Loader) convertWooormToTagSpecs(wooormData WooormAttributeData) []TagSpec {
	var specs []TagSpec
	processedElements := make(map[string]bool)

	globalAttrs := make(map[string]bool)
	if globals, exists := wooormData["*"]; exists {
		for _, attr := range globals {
			globalAttrs[attr] = true
		}
	}

	for tagName, attributes := range wooormData {
		if tagName == "*" {
			continue
		}

		spec := TagSpec{
			Name:          tagName,
			Void:          VoidElements[tagName],
			ParentTargets: ParentMap[tagName],
		}

		var elemAttributes []Attribute
		for _, attrName := range attributes {
			if attrName == "" || globalAttrs[attrName] {
				continue
			}

			field := camelCase(attrName)
			attr := Attribute{
				Field: field,
				Attr:  attrName,
				Type:  "string",
			}

			if BoolAttributes[strings.ToLower(attrName)] {
				attr.Type = "bool"
			}

			elemAttributes = append(elemAttributes, attr)
		}

		sort.Slice(elemAttributes, func(i, j int) bool {
			return elemAttributes[i].Attr < elemAttributes[j].Attr
		})

		spec.Attributes = elemAttributes
		specs = append(specs, spec)
		processedElements[tagName] = true
	}

	for elementName := range StandardHTML5Elements {
		if !processedElements[elementName] {
			spec := TagSpec{
				Name:          elementName,
				Void:          VoidElements[elementName],
				ParentTargets: ParentMap[elementName],
				Attributes:    []Attribute{},
			}
			specs = append(specs, spec)
		}
	}

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

	var keys []string
	for k := range elemData {
		if k == "__compat" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		field := camelCase(key)
		attr := Attribute{
			Field: field,
			Attr:  key,
			Type:  "string",
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

// CamelCase converts kebab-case to CamelCase
func CamelCase(name string) string {
	return camelCase(name)
}

func camelCase(name string) string {
	cleaned := strings.ReplaceAll(strings.ReplaceAll(name, "-", ""), "_", "")
	if len(cleaned) == 0 {
		return ""
	}
	return strings.ToUpper(cleaned[:1]) + cleaned[1:]
}
