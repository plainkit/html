package html

import "strings"

type FrameAttrs struct {
	Global       GlobalAttrs
	Frameborder  string
	Marginheight string
	Marginwidth  string
	Name         string
	Noresize     string
	Scrolling    string
	Src          string
}

type FrameArg interface {
	applyFrame(*FrameAttrs, *[]Component)
}

func defaultFrameAttrs() *FrameAttrs {
	return &FrameAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Frame(args ...FrameArg) Node {
	a := defaultFrameAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFrame(a, &kids)
	}
	return Node{Tag: "frame", Attrs: a, Kids: kids}
}

func (g Global) applyFrame(a *FrameAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o FrameborderOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Frameborder = o.v
}
func (o MarginheightOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Marginheight = o.v
}
func (o MarginwidthOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Marginwidth = o.v
}
func (o NameOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o NoresizeOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Noresize = o.v
}
func (o ScrollingOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Scrolling = o.v
}
func (o SrcOpt) applyFrame(a *FrameAttrs, _ *[]Component) {
	a.Src = o.v
}

func (a *FrameAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Frameborder != "" {
		Attr(sb, "frameborder", a.Frameborder)
	}
	if a.Marginheight != "" {
		Attr(sb, "marginheight", a.Marginheight)
	}
	if a.Marginwidth != "" {
		Attr(sb, "marginwidth", a.Marginwidth)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Noresize != "" {
		Attr(sb, "noresize", a.Noresize)
	}
	if a.Scrolling != "" {
		Attr(sb, "scrolling", a.Scrolling)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
}
