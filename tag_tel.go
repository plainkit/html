package html

import "strings"

type TelAttrs struct {
	Global GlobalAttrs
}

type TelArg interface {
	applyTel(*TelAttrs, *[]Component)
}

func defaultTelAttrs() *TelAttrs {
	return &TelAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tel(args ...TelArg) Node {
	a := defaultTelAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTel(a, &kids)
	}
	return Node{Tag: "tel", Attrs: a, Kids: kids}
}

func (g Global) applyTel(a *TelAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *TelAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
