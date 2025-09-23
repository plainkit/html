package html

import "strings"

type RangeAttrs struct {
	Global GlobalAttrs
}

type RangeArg interface {
	applyRange(*RangeAttrs, *[]Component)
}

func defaultRangeAttrs() *RangeAttrs {
	return &RangeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Range(args ...RangeArg) Node {
	a := defaultRangeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRange(a, &kids)
	}
	return Node{Tag: "range", Attrs: a, Kids: kids}
}

func (g Global) applyRange(a *RangeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyRange(_ *RangeAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyRange(_ *RangeAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *RangeAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
