package html

import "strings"

type FooterAttrs struct {
	Global GlobalAttrs
}

type FooterArg interface {
	applyFooter(*FooterAttrs, *[]Component)
}

func defaultFooterAttrs() *FooterAttrs {
	return &FooterAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Footer(args ...FooterArg) Node {
	a := defaultFooterAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFooter(a, &kids)
	}
	return Node{Tag: "footer", Attrs: a, Kids: kids}
}

func (g Global) applyFooter(a *FooterAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyFooter(_ *FooterAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyFooter(_ *FooterAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *FooterAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
