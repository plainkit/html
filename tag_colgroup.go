package html

import "strings"

type ColgroupAttrs struct {
	Global GlobalAttrs
	Span   string
}

type ColgroupArg interface {
	ApplyColgroup(*ColgroupAttrs, *[]Component)
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
		ar.ApplyColgroup(a, &kids)
	}
	return Node{Tag: "colgroup", Attrs: a, Kids: kids}
}

func (g Global) ApplyColgroup(a *ColgroupAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o SpanOpt) ApplyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Span = o.v
}

func (a *ColgroupAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Span != "" {
		Attr(sb, "span", a.Span)
	}
}
