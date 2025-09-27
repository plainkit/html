package html

import "strings"

type TheadAttrs struct {
	Global GlobalAttrs
}

type TheadArg interface {
	applyThead(*TheadAttrs, *[]Component)
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
		ar.applyThead(a, &kids)
	}
	return Node{Tag: "thead", Attrs: a, Kids: kids}
}

func (g Global) applyThead(a *TheadAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *TheadAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
