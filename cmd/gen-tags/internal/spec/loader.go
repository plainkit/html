package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

// ParentMap defines parent-child relationships (currently unused)
var ParentMap = map[string][]string{}

// Loader handles loading and parsing of HTML specification files
type Loader struct {
	specsDir string
}

// NewLoader creates a new spec loader for the given directory
func NewLoader(specsDir string) *Loader {
	return &Loader{specsDir: specsDir}
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

// camelCase converts kebab-case to CamelCase
func camelCase(name string) string {
	delimiters := func(r rune) bool { return r == '-' || r == '_' }
	parts := strings.FieldsFunc(name, delimiters)
	for i, p := range parts {
		if len(p) == 0 {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}
