package blox

import "strings"

// DD (Description Definition)
type DdAttrs struct {
	Global GlobalAttrs
}

type DdArg interface {
	applyDd(*DdAttrs, *[]Component)
}

func defaultDdAttrs() *DdAttrs {
	return &DdAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dd(args ...DdArg) Node {
	a := defaultDdAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDd(a, &kids)
	}
	return Node{Tag: "dd", Attrs: a, Kids: kids}
}

func (g Global) applyDd(a *DdAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDd(_ *DdAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDd(_ *DdAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *DdAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
