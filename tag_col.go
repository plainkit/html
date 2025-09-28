package html

import "strings"

type ColAttrs struct {
	Global GlobalAttrs
	Span   string
}

type ColArg interface {
	ApplyCol(*ColAttrs, *[]Component)
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
		ar.ApplyCol(a, &kids)
	}
	return Node{Tag: "col", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyCol(a *ColAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o SpanOpt) ApplyCol(a *ColAttrs, _ *[]Component) {
	a.Span = o.v
}

func (a *ColAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Span != "" {
		Attr(sb, "span", a.Span)
	}
}
