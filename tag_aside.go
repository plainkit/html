package html

import "strings"

type AsideAttrs struct {
	Global GlobalAttrs
}

type AsideArg interface {
	applyAside(*AsideAttrs, *[]Component)
}

func defaultAsideAttrs() *AsideAttrs {
	return &AsideAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Aside(args ...AsideArg) Node {
	a := defaultAsideAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAside(a, &kids)
	}
	return Node{Tag: "aside", Attrs: a, Kids: kids}
}

func (g Global) applyAside(a *AsideAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *AsideAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
