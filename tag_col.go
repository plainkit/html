package html

import "strings"

type ColAttrs struct {
	Global GlobalAttrs
	Span   string
}

type ColArg interface {
	applyCol(*ColAttrs, *[]Component)
}

func defaultColAttrs() *ColAttrs {
	return &ColAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Col(args ...ColArg) Node {
	a := defaultColAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCol(a, &kids)
	}
	return Node{Tag: "col", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyCol(a *ColAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o SpanOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Span = o.v
}

func (a *ColAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Span != "" {
		Attr(sb, "span", a.Span)
	}
}
