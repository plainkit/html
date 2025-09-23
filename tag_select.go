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
			Style:  "",
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

// Select-specific options
type SelectNameOpt struct{ v string }

func SelectName(v string) SelectNameOpt { return SelectNameOpt{v} }

func (g Global) applySelect(a *SelectAttrs, _ *[]Component)          { g.do(&a.Global) }
func (o SelectNameOpt) applySelect(a *SelectAttrs, _ *[]Component)   { a.Name = o.v }
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
