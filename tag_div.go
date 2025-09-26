package html

import "strings"

type DivAttrs struct {
	Global GlobalAttrs
	Align  string
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

func (o AlignOpt) applyDiv(a *DivAttrs, _ *[]Component) {
	a.Align = o.v
}

func (a *DivAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
