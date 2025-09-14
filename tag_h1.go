package blox

import "strings"

// H1
type H1Attrs struct {
	Global GlobalAttrs
}

type H1Arg interface {
	applyH1(*H1Attrs, *[]Component)
}

func defaultH1Attrs() *H1Attrs {
	return &H1Attrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H1(args ...H1Arg) Node {
	a := defaultH1Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH1(a, &kids)
	}
	return Node{Tag: "h1", Attrs: a, Kids: kids}
}

func (g Global) applyH1(a *H1Attrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyH1(_ *H1Attrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyH1(_ *H1Attrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *H1Attrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
