package blox

import "strings"

// Colgroup
type ColgroupAttrs struct {
	Global GlobalAttrs
	Span   int
}

type ColgroupArg interface {
	applyColgroup(*ColgroupAttrs, *[]Component)
}

func defaultColgroupAttrs() *ColgroupAttrs {
	return &ColgroupAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Colgroup(args ...ColgroupArg) Component {
	a := defaultColgroupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyColgroup(a, &kids)
	}
	return Node{Tag: "colgroup", Attrs: a, Kids: kids}
}

type SpanOpt struct{ v int }

func SpanAttr(v int) SpanOpt { return SpanOpt{v} }

func (g Global) applyColgroup(a *ColgroupAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyColgroup(_ *ColgroupAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyColgroup(_ *ColgroupAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o SpanOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component)     { a.Span = o.v }

func (a *ColgroupAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Span > 0 {
		attr(sb, "span", itoa(a.Span))
	}
}
