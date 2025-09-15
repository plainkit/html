package html

import "strings"

// H3
type H3Attrs struct {
	Global GlobalAttrs
}

type H3Arg interface {
	applyH3(*H3Attrs, *[]Component)
}

func defaultH3Attrs() *H3Attrs {
	return &H3Attrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H3(args ...H3Arg) Node {
	a := defaultH3Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH3(a, &kids)
	}
	return Node{Tag: "h3", Attrs: a, Kids: kids}
}

func (g Global) applyH3(a *H3Attrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyH3(_ *H3Attrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyH3(_ *H3Attrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *H3Attrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
