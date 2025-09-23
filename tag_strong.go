package html

import "strings"

type StrongAttrs struct {
	Global GlobalAttrs
}

type StrongArg interface {
	applyStrong(*StrongAttrs, *[]Component)
}

func defaultStrongAttrs() *StrongAttrs {
	return &StrongAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Strong(args ...StrongArg) Node {
	a := defaultStrongAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyStrong(a, &kids)
	}
	return Node{Tag: "strong", Attrs: a, Kids: kids}
}

func (g Global) applyStrong(a *StrongAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyStrong(_ *StrongAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyStrong(_ *StrongAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *StrongAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
