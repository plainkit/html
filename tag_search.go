package html

import "strings"

type SearchAttrs struct {
	Global GlobalAttrs
}

type SearchArg interface {
	ApplySearch(*SearchAttrs, *[]Component)
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
		ar.ApplySearch(a, &kids)
	}

	return Node{Tag: "search", Attrs: a, Kids: kids}
}

func (g Global) ApplySearch(a *SearchAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SearchAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
