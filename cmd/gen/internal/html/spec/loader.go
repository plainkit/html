package spec

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/plainkit/tags"
)

// Loader fetches HTML element specs from the local tags dataset.
type Loader struct {
	once    sync.Once
	loadErr error
	index   *tags.HTMLIndex
}

// NewLoader creates a loader backed by the tags dataset.
func NewLoader() *Loader {
	return &Loader{}
}

// LoadAllTagSpecs parses all tag definitions except the wildcard entry.
func (l *Loader) LoadAllTagSpecs() ([]TagSpec, error) {
	if err := l.ensureLoaded(); err != nil {
		return nil, err
	}

	globalSet := make(map[string]struct{}, len(l.index.Globals))
	for _, attr := range l.index.Globals {
		globalSet[strings.ToLower(attr.Name)] = struct{}{}
	}

	specs := make([]TagSpec, 0, len(l.index.Elements))
	for name, element := range l.index.Elements {
		if strings.TrimSpace(name) == "" {
			continue
		}

		tagAttrs := collectAttributes(element.Attributes, globalSet)
		applyTagOverrides(name, &tagAttrs)
		specs = append(specs, TagSpec{
			Name:       name,
			Void:       element.Empty,
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

	attrs := make([]Attribute, 0, len(l.index.Globals)+1)
	hasRole := false

	for _, attr := range l.index.Globals {
		key := strings.ToLower(attr.Name)

		attrs = append(attrs, Attribute{
			Field: CamelCase(key),
			Attr:  key,
			Type:  attributeTypeFromRef(attr),
		})
		if key == "role" {
			hasRole = true
		}
	}

	if !hasRole {
		attrs = append(attrs, Attribute{
			Field: CamelCase("role"),
			Attr:  "role",
			Type:  "string",
		})
	}

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
		l.index = &tags.HTML
		if l.index == nil {
			l.loadErr = fmt.Errorf("html index unavailable")
		}
	})

	return l.loadErr
}

func collectAttributes(attrRefs []tags.AttributeRef, globalSet map[string]struct{}) []Attribute {
	seen := make(map[string]struct{})
	attrs := make([]Attribute, 0, len(attrRefs))

	for _, ref := range attrRefs {
		key := strings.ToLower(ref.Name)
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
			Type:  attributeTypeFromRef(ref),
		})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Attr < attrs[j].Attr
	})

	return attrs
}

func applyTagOverrides(name string, attrs *[]Attribute) {
	if attrs == nil {
		return
	}

	ensureAttribute := func(field, attrType, attrName string) {
		for _, existing := range *attrs {
			if existing.Attr == attrName {
				return
			}
		}

		*attrs = append(*attrs, Attribute{Field: field, Type: attrType, Attr: attrName})
	}

	switch name {
	case "input":
		ensureAttribute("Name", "string", "name")
	}

	sort.Slice(*attrs, func(i, j int) bool {
		return (*attrs)[i].Attr < (*attrs)[j].Attr
	})
}

func attributeTypeFromRef(ref tags.AttributeRef) string {
	if ref.Boolean {
		return "bool"
	}

	return "string"
}
