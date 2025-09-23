package html

import "strings"

type TtAttrs struct {
	Global GlobalAttrs
}

type TtArg interface {
	applyTt(*TtAttrs, *[]Component)
}

func defaultTtAttrs() *TtAttrs {
	return &TtAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tt(args ...TtArg) Node {
	a := defaultTtAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTt(a, &kids)
	}
	return Node{Tag: "tt", Attrs: a, Kids: kids}
}

func (g Global) applyTt(a *TtAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTt(_ *TtAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTt(_ *TtAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *TtAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
