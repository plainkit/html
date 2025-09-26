package html

import "strings"

type FieldsetAttrs struct {
	Global   GlobalAttrs
	Disabled bool
	Form     string
	Name     string
}

type FieldsetArg interface {
	applyFieldset(*FieldsetAttrs, *[]Component)
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
		ar.applyFieldset(a, &kids)
	}
	return Node{Tag: "fieldset", Attrs: a, Kids: kids}
}

func (g Global) applyFieldset(a *FieldsetAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DisabledOpt) applyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FormOpt) applyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o NameOpt) applyFieldset(a *FieldsetAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *FieldsetAttrs) writeAttrs(sb *strings.Builder) {
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
