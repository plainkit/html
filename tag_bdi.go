package html

import "strings"

type BdiAttrs struct {
	Global GlobalAttrs
}

type BdiArg interface {
	ApplyBdi(*BdiAttrs, *[]Component)
}

func defaultBdiAttrs() *BdiAttrs {
	return &BdiAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Bdi(args ...BdiArg) Node {
	a := defaultBdiAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyBdi(a, &kids)
	}
	return Node{Tag: "bdi", Attrs: a, Kids: kids}
}

func (g Global) ApplyBdi(a *BdiAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *BdiAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
