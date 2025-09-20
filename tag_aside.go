package html

import "strings"

// Aside
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

func (g Global) applyAside(a *AsideAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyAside(_ *AsideAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyAside(_ *AsideAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *AsideAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
