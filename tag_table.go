package html

import "strings"

type TableAttrs struct {
	Global      GlobalAttrs
	Align       string
	Bgcolor     string
	Border      string
	Cellpadding string
	Cellspacing string
	Frame       string
	Rules       string
	Summary     string
	Width       string
}

type TableArg interface {
	applyTable(*TableAttrs, *[]Component)
}

func defaultTableAttrs() *TableAttrs {
	return &TableAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Table(args ...TableArg) Node {
	a := defaultTableAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTable(a, &kids)
	}
	return Node{Tag: "table", Attrs: a, Kids: kids}
}

func (g Global) applyTable(a *TableAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTable(_ *TableAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTable(_ *TableAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o AlignOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o BgcolorOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o BorderOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Border = o.v
}
func (o CellpaddingOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Cellpadding = o.v
}
func (o CellspacingOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Cellspacing = o.v
}
func (o FrameOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Frame = o.v
}
func (o RulesOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Rules = o.v
}
func (o SummaryOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Summary = o.v
}
func (o WidthOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *TableAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Bgcolor != "" {
		Attr(sb, "bgcolor", a.Bgcolor)
	}
	if a.Border != "" {
		Attr(sb, "border", a.Border)
	}
	if a.Cellpadding != "" {
		Attr(sb, "cellpadding", a.Cellpadding)
	}
	if a.Cellspacing != "" {
		Attr(sb, "cellspacing", a.Cellspacing)
	}
	if a.Frame != "" {
		Attr(sb, "frame", a.Frame)
	}
	if a.Rules != "" {
		Attr(sb, "rules", a.Rules)
	}
	if a.Summary != "" {
		Attr(sb, "summary", a.Summary)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
