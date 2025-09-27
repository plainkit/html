package html

import "strings"

type SpanAttrs struct {
	Global GlobalAttrs
}

type SpanArg interface {
	applySpan(*SpanAttrs, *[]Component)
}

func defaultSpanAttrs() *SpanAttrs {
	return &SpanAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Span(args ...SpanArg) Node {
	a := defaultSpanAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySpan(a, &kids)
	}
	return Node{Tag: "span", Attrs: a, Kids: kids}
}

func (g Global) applySpan(a *SpanAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *SpanAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
