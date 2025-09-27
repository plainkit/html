package spec

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/plainkit/html/cmd/gen/internal/svg/utils"
	"github.com/plainkit/tags"
)

type Loader struct {
	once    sync.Once
	loadErr error
	index   *tags.SVGIndex
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	if err := l.ensureLoaded(); err != nil {
		return nil, err
	}

	fmt.Println("Loading SVG element definitions from github.com/plainkit/tags...")

	globalSet := make(map[string]struct{}, len(l.index.Globals))
	for _, attr := range l.index.Globals {
		globalSet[strings.ToLower(attr)] = struct{}{}
	}

	htmlElements := l.htmlElementSet()

	specs := make([]TagSpec, 0, len(l.index.Elements))
	for name, element := range l.index.Elements {
		trimmed := strings.TrimSpace(name)
		if trimmed == "" {
			continue
		}

		lowerName := strings.ToLower(trimmed)
		if lowerName != "svg" {
			if _, exists := htmlElements[lowerName]; exists {
				continue
			}
		}

		attrs := collectAttributes(element.Attributes, globalSet)
		if lowerName == "svg" {
			attrs = append(attrs, Attribute{
				Field: "Xmlns",
				Attr:  "xmlns",
				Type:  "string",
			})
		}
		specs = append(specs, TagSpec{
			Name:       trimmed,
			Attributes: attrs,
		})
	}

	sort.Slice(specs, func(i, j int) bool {
		return specs[i].Name < specs[j].Name
	})

	fmt.Printf("Loaded %d SVG element definitions\n", len(specs))
	return specs, nil
}

func (l *Loader) LoadGlobalAttributes() ([]Attribute, error) {
	if err := l.ensureLoaded(); err != nil {
		return nil, err
	}

	htmlGlobals := l.htmlGlobalSet()

	attrs := make([]Attribute, 0, len(l.index.Globals))
	for _, name := range l.index.Globals {
		trimmed := strings.TrimSpace(name)
		if trimmed == "" {
			continue
		}
		lower := strings.ToLower(trimmed)
		if _, ok := htmlGlobals[lower]; ok {
			continue
		}

		attrs = append(attrs, Attribute{
			Field: utils.CamelCase(trimmed),
			Attr:  trimmed,
			Type:  "string",
		})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Attr < attrs[j].Attr
	})

	return attrs, nil
}

func (l *Loader) CollectAllAttributes(specs []TagSpec) map[string]Attribute {
	all := make(map[string]Attribute)

	for _, spec := range specs {
		for _, attr := range spec.Attributes {
			key := strings.ToLower(attr.Attr)
			if key == "" {
				continue
			}
			if existing, ok := all[key]; ok {
				if strings.Contains(attr.Attr, "-") && !strings.Contains(existing.Attr, "-") {
					all[key] = attr
				} else if existing.Type == "bool" && attr.Type == "string" {
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
		l.index = &tags.SVG
		if l.index == nil {
			l.loadErr = fmt.Errorf("svg index unavailable")
		}
	})

	return l.loadErr
}

func (l *Loader) htmlGlobalSet() map[string]struct{} {
	globals := make(map[string]struct{})

	for _, attr := range tags.HTML.Globals {
		name := strings.ToLower(attr.Name)
		if name != "" {
			globals[name] = struct{}{}
		}
	}

	return globals
}

func (l *Loader) htmlElementSet() map[string]struct{} {
	elements := make(map[string]struct{})

	for name := range tags.HTML.Elements {
		lower := strings.ToLower(name)
		if lower != "" {
			elements[lower] = struct{}{}
		}
	}

	return elements
}

func collectAttributes(attrNames []string, globalSet map[string]struct{}) []Attribute {
	seen := make(map[string]struct{})
	attrs := make([]Attribute, 0, len(attrNames))

	for _, attrName := range attrNames {
		trimmed := strings.TrimSpace(attrName)
		if trimmed == "" {
			continue
		}

		lower := strings.ToLower(trimmed)
		if _, ok := globalSet[lower]; ok {
			continue
		}
		if _, ok := seen[lower]; ok {
			continue
		}
		seen[lower] = struct{}{}

		attrs = append(attrs, Attribute{
			Field: utils.CamelCase(trimmed),
			Attr:  trimmed,
			Type:  "string",
		})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Attr < attrs[j].Attr
	})

	return attrs
}
