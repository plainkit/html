package html

import "strings"

type NumberAttrs struct {
	Global GlobalAttrs
}

type NumberArg interface {
	applyNumber(*NumberAttrs, *[]Component)
}

func defaultNumberAttrs() *NumberAttrs {
	return &NumberAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Number(args ...NumberArg) Node {
	a := defaultNumberAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNumber(a, &kids)
	}
	return Node{Tag: "number", Attrs: a, Kids: kids}
}

func (g Global) applyNumber(a *NumberAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NumberAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
