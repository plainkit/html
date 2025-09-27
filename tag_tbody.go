package html

import "strings"

type TbodyAttrs struct {
	Global  GlobalAttrs
	Align   string
	Char    string
	Charoff string
	Valign  string
}

type TbodyArg interface {
	applyTbody(*TbodyAttrs, *[]Component)
}

func defaultTbodyAttrs() *TbodyAttrs {
	return &TbodyAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tbody(args ...TbodyArg) Node {
	a := defaultTbodyAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTbody(a, &kids)
	}
	return Node{Tag: "tbody", Attrs: a, Kids: kids}
}

func (g Global) applyTbody(a *TbodyAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AlignOpt) applyTbody(a *TbodyAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o CharOpt) applyTbody(a *TbodyAttrs, _ *[]Component) {
	a.Char = o.v
}
func (o CharoffOpt) applyTbody(a *TbodyAttrs, _ *[]Component) {
	a.Charoff = o.v
}
func (o ValignOpt) applyTbody(a *TbodyAttrs, _ *[]Component) {
	a.Valign = o.v
}

func (a *TbodyAttrs) WriteAttrs(sb *strings.Builder) {
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
	if a.Valign != "" {
		Attr(sb, "valign", a.Valign)
	}
}
