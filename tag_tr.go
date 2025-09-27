package html

import "strings"

type TrAttrs struct {
	Global  GlobalAttrs
	Align   string
	Bgcolor string
	Char    string
	Charoff string
	Valign  string
}

type TrArg interface {
	applyTr(*TrAttrs, *[]Component)
}

func defaultTrAttrs() *TrAttrs {
	return &TrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tr(args ...TrArg) Node {
	a := defaultTrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTr(a, &kids)
	}
	return Node{Tag: "tr", Attrs: a, Kids: kids}
}

func (g Global) applyTr(a *TrAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AlignOpt) applyTr(a *TrAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o BgcolorOpt) applyTr(a *TrAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o CharOpt) applyTr(a *TrAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyTr(a *TrAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ValignOpt) applyTr(a *TrAttrs, _ *[]Component) {
	a.Valign = o.v
}

func (a *TrAttrs) WriteAttrs(sb *strings.Builder) {
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
