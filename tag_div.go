package html

import "strings"

type DivAttrs struct {
	Global GlobalAttrs
}

type DivArg interface {
	applyDiv(*DivAttrs, *[]Component)
}

func defaultDivAttrs() *DivAttrs {
	return &DivAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Div(args ...DivArg) Node {
	a := defaultDivAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDiv(a, &kids)
	}
	return Node{Tag: "div", Attrs: a, Kids: kids}
}

func (g Global) applyDiv(a *DivAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *DivAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
