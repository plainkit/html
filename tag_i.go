package html

import "strings"

type IAttrs struct {
	Global GlobalAttrs
}

type IArg interface {
	ApplyI(*IAttrs, *[]Component)
}

func defaultIAttrs() *IAttrs {
	return &IAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func I(args ...IArg) Node {
	a := defaultIAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyI(a, &kids)
	}
	return Node{Tag: "i", Attrs: a, Kids: kids}
}

func (g Global) ApplyI(a *IAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *IAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
