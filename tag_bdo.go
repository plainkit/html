package html

import "strings"

type BdoAttrs struct {
	Global GlobalAttrs
}

type BdoArg interface {
	ApplyBdo(*BdoAttrs, *[]Component)
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
		ar.ApplyBdo(a, &kids)
	}
	return Node{Tag: "bdo", Attrs: a, Kids: kids}
}

func (g Global) ApplyBdo(a *BdoAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *BdoAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
