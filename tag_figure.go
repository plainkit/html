package html

import "strings"

type FigureAttrs struct {
	Global GlobalAttrs
}

type FigureArg interface {
	applyFigure(*FigureAttrs, *[]Component)
}

func defaultFigureAttrs() *FigureAttrs {
	return &FigureAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Figure(args ...FigureArg) Node {
	a := defaultFigureAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFigure(a, &kids)
	}
	return Node{Tag: "figure", Attrs: a, Kids: kids}
}

func (g Global) applyFigure(a *FigureAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyFigure(_ *FigureAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyFigure(_ *FigureAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *FigureAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
