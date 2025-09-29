package html

import "strings"

type UAttrs struct {
	Global GlobalAttrs
}

type UArg interface {
	ApplyU(*UAttrs, *[]Component)
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
		ar.ApplyU(a, &kids)
	}

	return Node{Tag: "u", Attrs: a, Kids: kids}
}

func (g Global) ApplyU(a *UAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *UAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
