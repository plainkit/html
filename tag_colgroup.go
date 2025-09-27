package html

import "strings"

type ColgroupAttrs struct {
	Global GlobalAttrs
	Span   string
}

type ColgroupArg interface {
	applyColgroup(*ColgroupAttrs, *[]Component)
}

func defaultColgroupAttrs() *ColgroupAttrs {
	return &ColgroupAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Colgroup(args ...ColgroupArg) Node {
	a := defaultColgroupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyColgroup(a, &kids)
	}
	return Node{Tag: "colgroup", Attrs: a, Kids: kids}
}

func (g Global) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o SpanOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Span = o.v
}

func (a *ColgroupAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Span != "" {
		Attr(sb, "span", a.Span)
	}
}
