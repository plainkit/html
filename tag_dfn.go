package blox

import "strings"

// Dfn
type DfnAttrs struct {
	Global GlobalAttrs
}

type DfnArg interface {
	applyDfn(*DfnAttrs, *[]Component)
}

func defaultDfnAttrs() *DfnAttrs {
	return &DfnAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dfn(args ...DfnArg) Component {
	a := defaultDfnAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDfn(a, &kids)
	}
	return Node{Tag: "dfn", Attrs: a, Kids: kids}
}

func (g Global) applyDfn(a *DfnAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDfn(_ *DfnAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDfn(_ *DfnAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *DfnAttrs) writeAttrs(sb *strings.Builder)         { writeGlobal(sb, &a.Global) }
