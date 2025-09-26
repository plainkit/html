package html

import "strings"

type ColgroupAttrs struct {
	Global  GlobalAttrs
	Align   string
	Char    string
	Charoff string
	Span    string
	Valign  string
	Width   string
}

type ColgroupArg interface {
	applyColgroup(*ColgroupAttrs, *[]Component)
}

func defaultColgroupAttrs() *ColgroupAttrs {
	return &ColgroupAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Colgroup(args ...ColgroupArg) Node {
	a := defaultColgroupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyColgroup(a, &kids)
	}
	return Node{Tag: "colgroup", Attrs: a, Kids: kids}
}

func (g Global) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o CharOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o SpanOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Span = o.v
}
func (o ValignOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Valign = o.v
}
func (o WidthOpt) applyColgroup(a *ColgroupAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ColgroupAttrs) writeAttrs(sb *strings.Builder) {
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
