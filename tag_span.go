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
			Style:  map[string]string{},
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

// Global option glue
func (g Global) applySpan(a *SpanAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applySpan(_ *SpanAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applySpan(_ *SpanAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *SpanAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
