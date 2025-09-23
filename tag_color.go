package html

import "strings"

type ColorAttrs struct {
	Global GlobalAttrs
}

type ColorArg interface {
	applyColor(*ColorAttrs, *[]Component)
}

func defaultColorAttrs() *ColorAttrs {
	return &ColorAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Color(args ...ColorArg) Node {
	a := defaultColorAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyColor(a, &kids)
	}
	return Node{Tag: "color", Attrs: a, Kids: kids}
}

func (g Global) applyColor(a *ColorAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyColor(_ *ColorAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyColor(_ *ColorAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *ColorAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
