package html

import "strings"

type UlAttrs struct {
	Global GlobalAttrs
}

type UlArg interface {
	ApplyUl(*UlAttrs, *[]Component)
}

func defaultUlAttrs() *UlAttrs {
	return &UlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ul(args ...UlArg) Node {
	a := defaultUlAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyUl(a, &kids)
	}

	return Node{Tag: "ul", Attrs: a, Kids: kids}
}

func (g Global) ApplyUl(a *UlAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *UlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
