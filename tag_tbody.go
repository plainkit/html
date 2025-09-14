package blox

import "strings"

// Tbody
type TbodyAttrs struct {
	Global GlobalAttrs
}

type TbodyArg interface {
	applyTbody(*TbodyAttrs, *[]Component)
}

func defaultTbodyAttrs() *TbodyAttrs {
	return &TbodyAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tbody(args ...TbodyArg) Node {
	a := defaultTbodyAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTbody(a, &kids)
	}
	return Node{Tag: "tbody", Attrs: a, Kids: kids}
}

func (g Global) applyTbody(a *TbodyAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTbody(_ *TbodyAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTbody(_ *TbodyAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TbodyAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
