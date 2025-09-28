package html

import "strings"

type FigcaptionAttrs struct {
	Global GlobalAttrs
}

type FigcaptionArg interface {
	ApplyFigcaption(*FigcaptionAttrs, *[]Component)
}

func defaultFigcaptionAttrs() *FigcaptionAttrs {
	return &FigcaptionAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyFigcaption(a, &kids)
	}
	return Node{Tag: "figcaption", Attrs: a, Kids: kids}
}

func (g Global) ApplyFigcaption(a *FigcaptionAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *FigcaptionAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
