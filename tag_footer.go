package html

import "strings"

type FooterAttrs struct {
	Global GlobalAttrs
}

type FooterArg interface {
	ApplyFooter(*FooterAttrs, *[]Component)
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
		ar.ApplyFooter(a, &kids)
	}

	return Node{Tag: "footer", Attrs: a, Kids: kids}
}

func (g Global) ApplyFooter(a *FooterAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *FooterAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
