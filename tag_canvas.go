package html

import "strings"

type CanvasAttrs struct {
	Global    GlobalAttrs
	Height    string
	MozOpaque string
	Width     string
}

type CanvasArg interface {
	applyCanvas(*CanvasAttrs, *[]Component)
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
		ar.applyCanvas(a, &kids)
	}
	return Node{Tag: "canvas", Attrs: a, Kids: kids}
}

func (g Global) applyCanvas(a *CanvasAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o HeightOpt) applyCanvas(a *CanvasAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o MozOpaqueOpt) applyCanvas(a *CanvasAttrs, _ *[]Component) {
	a.MozOpaque = o.v
}
func (o WidthOpt) applyCanvas(a *CanvasAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *CanvasAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.MozOpaque != "" {
		Attr(sb, "moz-opaque", a.MozOpaque)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
