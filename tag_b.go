package html

import "strings"

type BAttrs struct {
	Global GlobalAttrs
}

type BArg interface {
	ApplyB(*BAttrs, *[]Component)
}

func defaultBAttrs() *BAttrs {
	return &BAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func B(args ...BArg) Node {
	a := defaultBAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyB(a, &kids)
	}

	return Node{Tag: "b", Attrs: a, Kids: kids}
}

func (g Global) ApplyB(a *BAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *BAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
