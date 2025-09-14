package blox

import "strings"

// Em
type EmAttrs struct {
	Global GlobalAttrs
}

type EmArg interface {
	applyEm(*EmAttrs, *[]Component)
}

func defaultEmAttrs() *EmAttrs {
	return &EmAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Em(args ...EmArg) Node {
	a := defaultEmAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyEm(a, &kids)
	}
	return Node{Tag: "em", Attrs: a, Kids: kids}
}

func (g Global) applyEm(a *EmAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyEm(_ *EmAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyEm(_ *EmAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *EmAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
