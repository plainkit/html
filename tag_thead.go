package html

import "strings"

type TheadAttrs struct {
	Global GlobalAttrs
}

type TheadArg interface {
	ApplyThead(*TheadAttrs, *[]Component)
}

func defaultTheadAttrs() *TheadAttrs {
	return &TheadAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Thead(args ...TheadArg) Node {
	a := defaultTheadAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyThead(a, &kids)
	}

	return Node{Tag: "thead", Attrs: a, Kids: kids}
}

func (g Global) ApplyThead(a *TheadAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *TheadAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
