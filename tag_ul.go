package html

import "strings"

// UL (Unordered List)
type UlAttrs struct {
	Global GlobalAttrs
}

type UlArg interface {
	applyUl(*UlAttrs, *[]Component)
}

func defaultUlAttrs() *UlAttrs {
	return &UlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ul(args ...UlArg) Node {
	a := defaultUlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyUl(a, &kids)
	}
	return Node{Tag: "ul", Attrs: a, Kids: kids}
}

func (g Global) applyUl(a *UlAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyUl(_ *UlAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyUl(_ *UlAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *UlAttrs) writeAttrs(sb *strings.Builder)        { writeGlobal(sb, &a.Global) }
