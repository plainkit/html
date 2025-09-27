package html

import "strings"

type FieldsetAttrs struct {
	Global GlobalAttrs
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

func (a *FieldsetAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
