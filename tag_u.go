package blox

import "strings"

// U
type UAttrs struct {
	Global GlobalAttrs
}

type UArg interface {
	applyU(*UAttrs, *[]Component)
}

func defaultUAttrs() *UAttrs {
	return &UAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func U(args ...UArg) Component {
	a := defaultUAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyU(a, &kids)
	}
	return Node{Tag: "u", Attrs: a, Kids: kids}
}

func (g Global) applyU(a *UAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyU(_ *UAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyU(_ *UAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *UAttrs) writeAttrs(sb *strings.Builder)       { writeGlobal(sb, &a.Global) }
