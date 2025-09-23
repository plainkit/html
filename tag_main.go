package html

import "strings"

type MainAttrs struct {
	Global GlobalAttrs
}

type MainArg interface {
	applyMain(*MainAttrs, *[]Component)
}

func defaultMainAttrs() *MainAttrs {
	return &MainAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Main(args ...MainArg) Node {
	a := defaultMainAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMain(a, &kids)
	}
	return Node{Tag: "main", Attrs: a, Kids: kids}
}

func (g Global) applyMain(a *MainAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyMain(_ *MainAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyMain(_ *MainAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *MainAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
