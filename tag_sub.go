package html

import "strings"

type SubAttrs struct {
	Global GlobalAttrs
}

type SubArg interface {
	applySub(*SubAttrs, *[]Component)
}

func defaultSubAttrs() *SubAttrs {
	return &SubAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Sub(args ...SubArg) Node {
	a := defaultSubAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySub(a, &kids)
	}
	return Node{Tag: "sub", Attrs: a, Kids: kids}
}

func (g Global) applySub(a *SubAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applySub(_ *SubAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applySub(_ *SubAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *SubAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
