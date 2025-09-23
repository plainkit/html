package html

import "strings"

type ColAttrs struct {
	Global  GlobalAttrs
	Align   string
	Char    string
	Charoff string
	Span    string
	Valign  string
	Width   string
}

type ColArg interface {
	applyCol(*ColAttrs, *[]Component)
}

func defaultColAttrs() *ColAttrs {
	return &ColAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Col(args ...ColArg) Node {
	a := defaultColAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCol(a, &kids)
	}
	return Node{Tag: "col", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyCol(a *ColAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyCol(_ *ColAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyCol(_ *ColAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o AlignOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o CharOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o SpanOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Span = o.v
}
func (o ValignOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Valign = o.v
}
func (o WidthOpt) applyCol(a *ColAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ColAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Char != "" {
		Attr(sb, "char", a.Char)
	}
	if a.Charoff != "" {
		Attr(sb, "charoff", a.Charoff)
	}
	if a.Span != "" {
		Attr(sb, "span", a.Span)
	}
	if a.Valign != "" {
		Attr(sb, "valign", a.Valign)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
