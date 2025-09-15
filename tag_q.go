package html

import "strings"

// Q
type QAttrs struct {
	Global GlobalAttrs
	Cite   string
}

type QArg interface {
	applyQ(*QAttrs, *[]Component)
}

func defaultQAttrs() *QAttrs {
	return &QAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Q(args ...QArg) Node {
	a := defaultQAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyQ(a, &kids)
	}
	return Node{Tag: "q", Attrs: a, Kids: kids}
}

func (g Global) applyQ(a *QAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyQ(_ *QAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyQ(_ *QAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o CiteOpt) applyQ(a *QAttrs, _ *[]Component)     { a.Cite = o.v }

func (a *QAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cite != "" {
		attr(sb, "cite", a.Cite)
	}
}
