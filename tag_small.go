package html

import "strings"

// Small
type SmallAttrs struct {
	Global GlobalAttrs
}

type SmallArg interface {
	applySmall(*SmallAttrs, *[]Component)
}

func defaultSmallAttrs() *SmallAttrs {
	return &SmallAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Small(args ...SmallArg) Node {
	a := defaultSmallAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySmall(a, &kids)
	}
	return Node{Tag: "small", Attrs: a, Kids: kids}
}

func (g Global) applySmall(a *SmallAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applySmall(_ *SmallAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applySmall(_ *SmallAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *SmallAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
