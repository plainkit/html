package spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
)

const htmlAttributesURL = "https://raw.githubusercontent.com/wooorm/html-element-attributes/refs/heads/main/index.js"

// Loader fetches and parses HTML element specs from wooorm/html-element-attributes.
type Loader struct {
	once     sync.Once
	loadErr  error
	elements map[string][]string
}

// NewLoader creates a loader that reads data from html-element-attributes.
func NewLoader() *Loader {
	return &Loader{}
}

// LoadAllTagSpecs parses all tag definitions except the wildcard entry.
func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	if err := l.ensureLoaded(); err != nil {
		return nil, err
	}

	globalSet := make(map[string]struct{})
	for _, attr := range l.elements["*"] {
		globalSet[strings.ToLower(attr)] = struct{}{}
	}

	var specs []TagSpec
	for name, attrList := range l.elements {
		if name == "*" {
			continue
		}

		if skipElements[name] {
			continue
		}

		tagAttrs := collectAttributes(attrList, globalSet)
		specs = append(specs, TagSpec{
			Name:       name,
			Void:       voidElements[name],
			Attributes: tagAttrs,
		})
	}

	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	return specs, nil
}

// LoadGlobalAttributes returns the wildcard attributes as global attributes.
func (l *Loader) LoadGlobalAttributes() ([]Attribute, error) {
	if err := l.ensureLoaded(); err != nil {
		return nil, err
	}

	attrs := collectAttributes(l.elements["*"], nil)

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Attr < attrs[j].Attr
	})

	return attrs, nil
}

// CollectAllAttributes aggregates attributes from tag specs (excluding globals).
func (l *Loader) CollectAllAttributes(specs []TagSpec) map[string]Attribute {
	all := make(map[string]Attribute)

	for _, spec := range specs {
		for _, attr := range spec.Attributes {
			key := strings.ToLower(attr.Attr)
			if existing, ok := all[key]; ok {
				if existing.Type == "bool" && attr.Type == "string" {
					all[key] = attr
				}
				continue
			}
			all[key] = attr
		}
	}

	return all
}

func (l *Loader) ensureLoaded() error {
	l.once.Do(func() {
		data, err := fetchSpec()
		if err != nil {
			l.loadErr = err
			return
		}

		attrs, err := parseSpec(data)
		if err != nil {
			l.loadErr = err
			return
		}

		l.elements = attrs
	})

	return l.loadErr
}

func fetchSpec() ([]byte, error) {
	resp, err := http.Get(htmlAttributesURL) // #nosec G107
	if err != nil {
		return nil, fmt.Errorf("fetch html-element-attributes: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	return body, nil
}

func parseSpec(data []byte) (map[string][]string, error) {
	const prefix = "export const htmlElementAttributes ="
	idx := bytes.Index(data, []byte(prefix))
	if idx == -1 {
		return nil, fmt.Errorf("unexpected spec format: missing export")
	}

	slice := data[idx+len(prefix):]
	start := bytes.IndexByte(slice, '{')
	end := bytes.LastIndexByte(slice, '}')
	if start == -1 || end == -1 || end <= start {
		return nil, fmt.Errorf("unexpected spec format")
	}

	raw := string(slice[start : end+1])

	reKey := regexp.MustCompile(`(?m)^\s*([A-Za-z0-9_-]+)\s*:`)
	raw = reKey.ReplaceAllString(raw, `"$1":`)
	raw = strings.ReplaceAll(raw, "'", "\"")

	var parsed map[string][]string
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return nil, fmt.Errorf("unmarshal spec JSON: %w", err)
	}

	return parsed, nil
}

func collectAttributes(attrNames []string, globalSet map[string]struct{}) []Attribute {
	seen := make(map[string]struct{})
	var attrs []Attribute

	for _, name := range attrNames {
		key := strings.ToLower(name)
		if key == "" {
			continue
		}
		if globalSet != nil {
			if _, ok := globalSet[key]; ok {
				continue
			}
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}

		attrs = append(attrs, Attribute{
			Field: CamelCase(key),
			Attr:  key,
			Type:  attributeType(key),
		})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Attr < attrs[j].Attr
	})

	return attrs
}

func attributeType(name string) string {
	name = strings.ToLower(name)
	if _, ok := booleanAttributes[name]; ok {
		return "bool"
	}
	return "string"
}

var booleanAttributes = map[string]struct{}{
	"allowfullscreen":     {},
	"allowpaymentrequest": {},
	"async":               {},
	"autofocus":           {},
	"autoplay":            {},
	"compact":             {},
	"checked":             {},
	"controls":            {},
	"default":             {},
	"declare":             {},
	"defer":               {},
	"disabled":            {},
	"formnovalidate":      {},
	"hidden":              {},
	"inert":               {},
	"ismap":               {},
	"itemscope":           {},
	"loop":                {},
	"multiple":            {},
	"muted":               {},
	"nohref":              {},
	"nomodule":            {},
	"noresize":            {},
	"noshade":             {},
	"novalidate":          {},
	"open":                {},
	"playsinline":         {},
	"readonly":            {},
	"required":            {},
	"reversed":            {},
	"selected":            {},
	"typemustmatch":       {},
	"nowrap":              {},
}

var voidElements = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"keygen": true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

var skipElements = map[string]bool{
	"":    true,
	"svg": true,
}
