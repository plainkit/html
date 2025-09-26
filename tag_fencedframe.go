package html

import "strings"

type FencedframeAttrs struct {
	Global GlobalAttrs
	Allow  string
	Height string
	Width  string
}

type FencedframeArg interface {
	applyFencedframe(*FencedframeAttrs, *[]Component)
}

func defaultFencedframeAttrs() *FencedframeAttrs {
	return &FencedframeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Fencedframe(args ...FencedframeArg) Node {
	a := defaultFencedframeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFencedframe(a, &kids)
	}
	return Node{Tag: "fencedframe", Attrs: a, Kids: kids}
}

func (g Global) applyFencedframe(a *FencedframeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AllowOpt) applyFencedframe(a *FencedframeAttrs, _ *[]Component) {
	a.Allow = o.v
}
func (o HeightOpt) applyFencedframe(a *FencedframeAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o WidthOpt) applyFencedframe(a *FencedframeAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *FencedframeAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Allow != "" {
		Attr(sb, "allow", a.Allow)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
