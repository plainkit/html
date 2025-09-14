package blox

import "strings"

// Sup
type SupAttrs struct {
	Global GlobalAttrs
}

type SupArg interface {
	applySup(*SupAttrs, *[]Component)
}

func defaultSupAttrs() *SupAttrs {
	return &SupAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Sup(args ...SupArg) Node {
	a := defaultSupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySup(a, &kids)
	}
	return Node{Tag: "sup", Attrs: a, Kids: kids}
}

func (g Global) applySup(a *SupAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applySup(_ *SupAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applySup(_ *SupAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *SupAttrs) writeAttrs(sb *strings.Builder)         { writeGlobal(sb, &a.Global) }
