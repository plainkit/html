package html

import "strings"

type UAttrs struct {
	Global GlobalAttrs
}

type UArg interface {
	applyU(*UAttrs, *[]Component)
}

func defaultUAttrs() *UAttrs {
	return &UAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func U(args ...UArg) Node {
	a := defaultUAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyU(a, &kids)
	}
	return Node{Tag: "u", Attrs: a, Kids: kids}
}

func (g Global) applyU(a *UAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *UAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
