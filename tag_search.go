package html

import "strings"

type SearchAttrs struct {
	Global GlobalAttrs
}

type SearchArg interface {
	applySearch(*SearchAttrs, *[]Component)
}

func defaultSearchAttrs() *SearchAttrs {
	return &SearchAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Search(args ...SearchArg) Node {
	a := defaultSearchAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySearch(a, &kids)
	}
	return Node{Tag: "search", Attrs: a, Kids: kids}
}

func (g Global) applySearch(a *SearchAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applySearch(_ *SearchAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applySearch(_ *SearchAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *SearchAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
