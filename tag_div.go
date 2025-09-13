package blox

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
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Div(args ...DivArg) Component {
	a := defaultDivAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDiv(a, &kids)
	}
	return Node{Tag: "div", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyDiv(a *DivAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyDiv(_ *DivAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyDiv(_ *DivAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *DivAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
