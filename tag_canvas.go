package html

import "strings"

// Canvas
type CanvasAttrs struct {
	Global GlobalAttrs
	Width  int
	Height int
}

type CanvasArg interface {
	applyCanvas(*CanvasAttrs, *[]Component)
}

func defaultCanvasAttrs() *CanvasAttrs {
	return &CanvasAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
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

func (g Global) applyCanvas(a *CanvasAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyCanvas(_ *CanvasAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyCanvas(_ *CanvasAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o WidthOpt) applyCanvas(a *CanvasAttrs, _ *[]Component)    { a.Width = o.v }
func (o HeightOpt) applyCanvas(a *CanvasAttrs, _ *[]Component)   { a.Height = o.v }

func (a *CanvasAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Width > 0 {
		attr(sb, "width", itoa(a.Width))
	}
	if a.Height > 0 {
		attr(sb, "height", itoa(a.Height))
	}
}
