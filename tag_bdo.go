package html

import "strings"

type BdoAttrs struct {
	Global GlobalAttrs
}

type BdoArg interface {
	applyBdo(*BdoAttrs, *[]Component)
}

func defaultBdoAttrs() *BdoAttrs {
	return &BdoAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Bdo(args ...BdoArg) Node {
	a := defaultBdoAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBdo(a, &kids)
	}
	return Node{Tag: "bdo", Attrs: a, Kids: kids}
}

func (g Global) applyBdo(a *BdoAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *BdoAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
