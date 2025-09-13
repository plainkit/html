package blox

import "strings"

// Fieldset
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
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Fieldset(args ...FieldsetArg) Component {
	a := defaultFieldsetAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFieldset(a, &kids)
	}
	return Node{Tag: "fieldset", Attrs: a, Kids: kids}
}

func (g Global) applyFieldset(a *FieldsetAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyFieldset(_ *FieldsetAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyFieldset(_ *FieldsetAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o DisabledOpt) applyFieldset(a *FieldsetAttrs, _ *[]Component) { a.Disabled = true }
func (o FormOpt) applyFieldset(a *FieldsetAttrs, _ *[]Component)     { a.Form = o.v }

func (a *FieldsetAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Disabled {
		boolAttr(sb, "disabled")
	}
	if a.Form != "" {
		attr(sb, "form", a.Form)
	}
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
}
