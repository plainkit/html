package html

import "strings"

// Select
type SelectAttrs struct {
	Global       GlobalAttrs
	Name         string
	Multiple     bool
	Required     bool
	Disabled     bool
	Size         int
	Form         string
	Autocomplete string
}

type SelectArg interface {
	applySelect(*SelectAttrs, *[]Component)
}

func defaultSelectAttrs() *SelectAttrs {
	return &SelectAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Select(args ...SelectArg) Node {
	a := defaultSelectAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySelect(a, &kids)
	}
	return Node{Tag: "select", Attrs: a, Kids: kids}
}

func (g Global) applySelect(a *SelectAttrs, _ *[]Component)          { g.do(&a.Global) }
func (o TxtOpt) applySelect(_ *SelectAttrs, kids *[]Component)       { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applySelect(_ *SelectAttrs, kids *[]Component)     { *kids = append(*kids, o.c) }
func (o MultipleOpt) applySelect(a *SelectAttrs, _ *[]Component)     { a.Multiple = true }
func (o RequiredOpt) applySelect(a *SelectAttrs, _ *[]Component)     { a.Required = true }
func (o DisabledOpt) applySelect(a *SelectAttrs, _ *[]Component)     { a.Disabled = true }
func (o SizeOpt) applySelect(a *SelectAttrs, _ *[]Component)         { a.Size = o.v }
func (o FormOpt) applySelect(a *SelectAttrs, _ *[]Component)         { a.Form = o.v }
func (o AutocompleteOpt) applySelect(a *SelectAttrs, _ *[]Component) { a.Autocomplete = o.v }

func (a *SelectAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
	if a.Multiple {
		boolAttr(sb, "multiple")
	}
	if a.Required {
		boolAttr(sb, "required")
	}
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
	if a.Size > 0 {
		attr(sb, "size", itoa(a.Size))
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
	if a.Autocomplete != "" {
		attr(sb, "autocomplete", a.Autocomplete)
	}
}
