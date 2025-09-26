package html

import "strings"

type TheadAttrs struct {
	Global  GlobalAttrs
	Align   string
	Bgcolor string
	Char    string
	Charoff string
	Valign  string
}

type TheadArg interface {
	applyThead(*TheadAttrs, *[]Component)
}

func defaultTheadAttrs() *TheadAttrs {
	return &TheadAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Thead(args ...TheadArg) Node {
	a := defaultTheadAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyThead(a, &kids)
	}
	return Node{Tag: "thead", Attrs: a, Kids: kids}
}

func (g Global) applyThead(a *TheadAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyThead(a *TheadAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o BgcolorOpt) applyThead(a *TheadAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o CharOpt) applyThead(a *TheadAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyThead(a *TheadAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ValignOpt) applyThead(a *TheadAttrs, _ *[]Component) {
	a.Valign = o.v
}

func (a *TheadAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
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
	if a.Valign != "" {
		Attr(sb, "valign", a.Valign)
	}
}
