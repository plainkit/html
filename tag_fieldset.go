package html

import "strings"

type FieldsetAttrs struct {
	Global   GlobalAttrs
	Disabled bool
	Form     string
	Name     string
}

type FieldsetArg interface {
	ApplyFieldset(*FieldsetAttrs, *[]Component)
}

func defaultFieldsetAttrs() *FieldsetAttrs {
	return &FieldsetAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Fieldset(args ...FieldsetArg) Node {
	a := defaultFieldsetAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyFieldset(a, &kids)
	}

	return Node{Tag: "fieldset", Attrs: a, Kids: kids}
}

func (g Global) ApplyFieldset(a *FieldsetAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o DisabledOpt) ApplyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) ApplyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o NameOpt) ApplyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *FieldsetAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Disabled {
		BoolAttr(sb, "disabled")
	}

	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}

	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
