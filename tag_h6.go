package html

import "strings"

// H6
type H6Attrs struct {
	Global GlobalAttrs
}

type H6Arg interface {
	applyH6(*H6Attrs, *[]Component)
}

func defaultH6Attrs() *H6Attrs {
	return &H6Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H6(args ...H6Arg) Node {
	a := defaultH6Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH6(a, &kids)
	}
	return Node{Tag: "h6", Attrs: a, Kids: kids}
}

func (g Global) applyH6(a *H6Attrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyH6(_ *H6Attrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyH6(_ *H6Attrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *H6Attrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
