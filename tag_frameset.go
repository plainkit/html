package html

import "strings"

type FramesetAttrs struct {
	Global GlobalAttrs
	Cols   string
	Rows   string
}

type FramesetArg interface {
	applyFrameset(*FramesetAttrs, *[]Component)
}

func defaultFramesetAttrs() *FramesetAttrs {
	return &FramesetAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Frameset(args ...FramesetArg) Node {
	a := defaultFramesetAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFrameset(a, &kids)
	}
	return Node{Tag: "frameset", Attrs: a, Kids: kids}
}

func (g Global) applyFrameset(a *FramesetAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyFrameset(_ *FramesetAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyFrameset(_ *FramesetAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o ColsOpt) applyFrameset(a *FramesetAttrs, _ *[]Component) {
	a.Cols = o.v
}
func (o RowsOpt) applyFrameset(a *FramesetAttrs, _ *[]Component) {
	a.Rows = o.v
}

func (a *FramesetAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Cols != "" {
		Attr(sb, "cols", a.Cols)
	}
	if a.Rows != "" {
		Attr(sb, "rows", a.Rows)
	}
}
