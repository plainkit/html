package html

import "strings"

type AbbrAttrs struct {
	Global GlobalAttrs
}

type AbbrArg interface {
	applyAbbr(*AbbrAttrs, *[]Component)
}

func defaultAbbrAttrs() *AbbrAttrs {
	return &AbbrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Abbr(args ...AbbrArg) Node {
	a := defaultAbbrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAbbr(a, &kids)
	}
	return Node{Tag: "abbr", Attrs: a, Kids: kids}
}

func (g Global) applyAbbr(a *AbbrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyAbbr(_ *AbbrAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyAbbr(_ *AbbrAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *AbbrAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
