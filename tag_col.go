package html

import "strings"

// Col (void)
type ColAttrs struct {
	Global GlobalAttrs
	Span   int
}

type ColArg interface {
	applyCol(*ColAttrs)
}

func defaultColAttrs() *ColAttrs {
	return &ColAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Col(args ...ColArg) Node {
	a := defaultColAttrs()
	for _, ar := range args {
		ar.applyCol(a)
	}
	return Node{Tag: "col", Attrs: a, Void: true}
}

func (g Global) applyCol(a *ColAttrs)  { g.do(&a.Global) }
func (o SpanOpt) applyCol(a *ColAttrs) { a.Span = o.v }

func (a *ColAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Span > 0 {
		attr(sb, "span", itoa(a.Span))
	}
}
