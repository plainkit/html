package html

import "strings"

type NobrAttrs struct {
	Global GlobalAttrs
}

type NobrArg interface {
	applyNobr(*NobrAttrs, *[]Component)
}

func defaultNobrAttrs() *NobrAttrs {
	return &NobrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Nobr(args ...NobrArg) Node {
	a := defaultNobrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNobr(a, &kids)
	}
	return Node{Tag: "nobr", Attrs: a, Kids: kids}
}

func (g Global) applyNobr(a *NobrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NobrAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
