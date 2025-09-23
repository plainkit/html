package html

import "strings"

type ThAttrs struct {
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

type ThArg interface {
	applyTh(*ThAttrs, *[]Component)
}

func defaultThAttrs() *ThAttrs {
	return &ThAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Th(args ...ThArg) Node {
	a := defaultThAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTh(a, &kids)
	}
	return Node{Tag: "th", Attrs: a, Kids: kids}
}

func (g Global) applyTh(a *ThAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTh(_ *ThAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTh(_ *ThAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o AbbrOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Abbr = o.v
}
func (o AlignOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o AxisOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Axis = o.v
}
func (o BgcolorOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o CharOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ColspanOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Colspan = o.v
}
func (o HeadersOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Headers = o.v
}
func (o RowspanOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Rowspan = o.v
}
func (o ScopeOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Scope = o.v
}
func (o ValignOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Valign = o.v
}
func (o WidthOpt) applyTh(a *ThAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ThAttrs) writeAttrs(sb *strings.Builder) {
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
