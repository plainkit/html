package html

import "strings"

// Figcaption
type FigcaptionAttrs struct {
	Global GlobalAttrs
}

type FigcaptionArg interface {
	applyFigcaption(*FigcaptionAttrs, *[]Component)
}

func defaultFigcaptionAttrs() *FigcaptionAttrs {
	return &FigcaptionAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Figcaption(args ...FigcaptionArg) Node {
	a := defaultFigcaptionAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFigcaption(a, &kids)
	}
	return Node{Tag: "figcaption", Attrs: a, Kids: kids}
}

func (g Global) applyFigcaption(a *FigcaptionAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyFigcaption(_ *FigcaptionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyFigcaption(_ *FigcaptionAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *FigcaptionAttrs) writeAttrs(sb *strings.Builder)                { writeGlobal(sb, &a.Global) }
