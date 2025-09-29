package html

import "strings"

type SelectAttrs struct {
	Global       GlobalAttrs
	Autocomplete string
	Disabled     bool
	Form         string
	Multiple     bool
	Name         string
	Required     bool
	Size         string
}

type SelectArg interface {
	ApplySelect(*SelectAttrs, *[]Component)
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
		ar.ApplySelect(a, &kids)
	}

	return Node{Tag: "select", Attrs: a, Kids: kids}
}

func (g Global) ApplySelect(a *SelectAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AutocompleteOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Autocomplete = o.v
}
func (o DisabledOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o MultipleOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Multiple = true
}
func (o NameOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o RequiredOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Required = true
}
func (o SizeOpt) ApplySelect(a *SelectAttrs, _ *[]Component) {
	a.Size = o.v
}

func (a *SelectAttrs) WriteAttrs(sb *strings.Builder) {
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
