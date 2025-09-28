package html

import "strings"

type HrAttrs struct {
	Global GlobalAttrs
}

type HrArg interface {
	ApplyHr(*HrAttrs, *[]Component)
}

func defaultHrAttrs() *HrAttrs {
	return &HrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Hr(args ...HrArg) Node {
	a := defaultHrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyHr(a, &kids)
	}
	return Node{Tag: "hr", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyHr(a *HrAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *HrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
