package html

import "strings"

type DdAttrs struct {
	Global GlobalAttrs
}

type DdArg interface {
	applyDd(*DdAttrs, *[]Component)
}

func defaultDdAttrs() *DdAttrs {
	return &DdAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dd(args ...DdArg) Node {
	a := defaultDdAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDd(a, &kids)
	}
	return Node{Tag: "dd", Attrs: a, Kids: kids}
}

func (g Global) applyDd(a *DdAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *DdAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
