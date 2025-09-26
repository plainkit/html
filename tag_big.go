package html

import "strings"

type BigAttrs struct {
	Global GlobalAttrs
}

type BigArg interface {
	applyBig(*BigAttrs, *[]Component)
}

func defaultBigAttrs() *BigAttrs {
	return &BigAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Big(args ...BigArg) Node {
	a := defaultBigAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBig(a, &kids)
	}
	return Node{Tag: "big", Attrs: a, Kids: kids}
}

func (g Global) applyBig(a *BigAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *BigAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
