package html

import "strings"

type RpAttrs struct {
	Global GlobalAttrs
}

type RpArg interface {
	applyRp(*RpAttrs, *[]Component)
}

func defaultRpAttrs() *RpAttrs {
	return &RpAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rp(args ...RpArg) Node {
	a := defaultRpAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRp(a, &kids)
	}
	return Node{Tag: "rp", Attrs: a, Kids: kids}
}

func (g Global) applyRp(a *RpAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyRp(_ *RpAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyRp(_ *RpAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *RpAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
