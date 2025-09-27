package html

import "strings"

type HrAttrs struct {
	Global GlobalAttrs
}

type HrArg interface {
	applyHr(*HrAttrs, *[]Component)
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
		ar.applyHr(a, &kids)
	}
	return Node{Tag: "hr", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyHr(a *HrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *HrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
