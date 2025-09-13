package blox

import "strings"

// Tr
type TrAttrs struct {
	Global GlobalAttrs
}

type TrArg interface {
	applyTr(*TrAttrs, *[]Component)
}

func defaultTrAttrs() *TrAttrs {
	return &TrAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tr(args ...TrArg) Component {
	a := defaultTrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTr(a, &kids)
	}
	return Node{Tag: "tr", Attrs: a, Kids: kids}
}

func (g Global) applyTr(a *TrAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTr(_ *TrAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTr(_ *TrAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TrAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
