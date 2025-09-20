package html

import "strings"

// B
type BAttrs struct {
	Global GlobalAttrs
}

type BArg interface {
	applyB(*BAttrs, *[]Component)
}

func defaultBAttrs() *BAttrs {
	return &BAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func B(args ...BArg) Node {
	a := defaultBAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyB(a, &kids)
	}
	return Node{Tag: "b", Attrs: a, Kids: kids}
}

func (g Global) applyB(a *BAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyB(_ *BAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyB(_ *BAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *BAttrs) writeAttrs(sb *strings.Builder)       { writeGlobal(sb, &a.Global) }
