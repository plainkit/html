package html

import "strings"

type TdAttrs struct {
	Global  GlobalAttrs
	Abbr    string
	Align   string
	Axis    string
	Bgcolor string
	Char    string
	Charoff string
	Colspan string
	Headers string
	Rowspan string
	Scope   string
	Valign  string
	Width   string
}

type TdArg interface {
	applyTd(*TdAttrs, *[]Component)
}

func defaultTdAttrs() *TdAttrs {
	return &TdAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Td(args ...TdArg) Node {
	a := defaultTdAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTd(a, &kids)
	}
	return Node{Tag: "td", Attrs: a, Kids: kids}
}

func (g Global) applyTd(a *TdAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTd(_ *TdAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTd(_ *TdAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o AbbrOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Abbr = o.v
}
func (o AlignOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o AxisOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Axis = o.v
}
func (o BgcolorOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o CharOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ColspanOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Colspan = o.v
}
func (o HeadersOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Headers = o.v
}
func (o RowspanOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Rowspan = o.v
}
func (o ScopeOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Scope = o.v
}
func (o ValignOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Valign = o.v
}
func (o WidthOpt) applyTd(a *TdAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *TdAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Abbr != "" {
		Attr(sb, "abbr", a.Abbr)
	}
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Axis != "" {
		Attr(sb, "axis", a.Axis)
	}
	if a.Bgcolor != "" {
		Attr(sb, "bgcolor", a.Bgcolor)
	}
	if a.Char != "" {
		Attr(sb, "char", a.Char)
	}
	if a.Charoff != "" {
		Attr(sb, "charoff", a.Charoff)
	}
	if a.Colspan != "" {
		Attr(sb, "colspan", a.Colspan)
	}
	if a.Headers != "" {
		Attr(sb, "headers", a.Headers)
	}
	if a.Rowspan != "" {
		Attr(sb, "rowspan", a.Rowspan)
	}
	if a.Scope != "" {
		Attr(sb, "scope", a.Scope)
	}
	if a.Valign != "" {
		Attr(sb, "valign", a.Valign)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
