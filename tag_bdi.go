package html

import "strings"

type BdiAttrs struct {
	Global GlobalAttrs
}

type BdiArg interface {
	applyBdi(*BdiAttrs, *[]Component)
}

func defaultBdiAttrs() *BdiAttrs {
	return &BdiAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Bdi(args ...BdiArg) Node {
	a := defaultBdiAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBdi(a, &kids)
	}
	return Node{Tag: "bdi", Attrs: a, Kids: kids}
}

func (g Global) applyBdi(a *BdiAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyBdi(_ *BdiAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyBdi(_ *BdiAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *BdiAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
