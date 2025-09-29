package html

import "strings"

type TrAttrs struct {
	Global GlobalAttrs
}

type TrArg interface {
	ApplyTr(*TrAttrs, *[]Component)
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
		ar.ApplyTr(a, &kids)
	}

	return Node{Tag: "tr", Attrs: a, Kids: kids}
}

func (g Global) ApplyTr(a *TrAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *TrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
