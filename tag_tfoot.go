package html

import "strings"

type TfootAttrs struct {
	Global  GlobalAttrs
	Align   string
	Bgcolor string
	Char    string
	Charoff string
	Valign  string
}

type TfootArg interface {
	applyTfoot(*TfootAttrs, *[]Component)
}

func defaultTfootAttrs() *TfootAttrs {
	return &TfootAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tfoot(args ...TfootArg) Node {
	a := defaultTfootAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTfoot(a, &kids)
	}
	return Node{Tag: "tfoot", Attrs: a, Kids: kids}
}

func (g Global) applyTfoot(a *TfootAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTfoot(_ *TfootAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTfoot(_ *TfootAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o AlignOpt) applyTfoot(a *TfootAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o BgcolorOpt) applyTfoot(a *TfootAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o CharOpt) applyTfoot(a *TfootAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyTfoot(a *TfootAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ValignOpt) applyTfoot(a *TfootAttrs, _ *[]Component) {
	a.Valign = o.v
}

func (a *TfootAttrs) writeAttrs(sb *strings.Builder) {
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
