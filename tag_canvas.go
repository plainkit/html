package html

import "strings"

type CanvasAttrs struct {
	Global GlobalAttrs
	Height string
	Width  string
}

type CanvasArg interface {
	ApplyCanvas(*CanvasAttrs, *[]Component)
}

func defaultCanvasAttrs() *CanvasAttrs {
	return &CanvasAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Canvas(args ...CanvasArg) Node {
	a := defaultCanvasAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyCanvas(a, &kids)
	}
	return Node{Tag: "canvas", Attrs: a, Kids: kids}
}

func (g Global) ApplyCanvas(a *CanvasAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o HeightOpt) ApplyCanvas(a *CanvasAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o WidthOpt) ApplyCanvas(a *CanvasAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *CanvasAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
