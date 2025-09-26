package html

import "strings"

type SelectAttrs struct {
	Global       GlobalAttrs
	Autocomplete string
	Disabled     bool
	Form         string
	HrInSelect   string
	Multiple     bool
	Name         string
	Required     bool
	Size         string
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

func (g Global) applySelect(a *SelectAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AutocompleteOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Autocomplete = o.v
}
func (o DisabledOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o HrInSelectOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.HrInSelect = o.v
}
func (o MultipleOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Multiple = true
}
func (o NameOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o RequiredOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Required = true
}
func (o SizeOpt) applySelect(a *SelectAttrs, _ *[]Component) {
	a.Size = o.v
}

func (a *SelectAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Autocomplete != "" {
		Attr(sb, "autocomplete", a.Autocomplete)
	}
	if a.Disabled {
		BoolAttr(sb, "disabled")
	}
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.HrInSelect != "" {
		Attr(sb, "hr_in_select", a.HrInSelect)
	}
	if a.Multiple {
		BoolAttr(sb, "multiple")
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Required {
		BoolAttr(sb, "required")
	}
	if a.Size != "" {
		Attr(sb, "size", a.Size)
	}
}
