package blox

import "strings"

// I
type IAttrs struct {
	Global GlobalAttrs
}

type IArg interface {
	applyI(*IAttrs, *[]Component)
}

func defaultIAttrs() *IAttrs {
	return &IAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func I(args ...IArg) Component {
	a := defaultIAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyI(a, &kids)
	}
	return Node{Tag: "i", Attrs: a, Kids: kids}
}

func (g Global) applyI(a *IAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyI(_ *IAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyI(_ *IAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *IAttrs) writeAttrs(sb *strings.Builder)       { writeGlobal(sb, &a.Global) }
