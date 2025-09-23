package html

import "strings"

type MarkAttrs struct {
	Global GlobalAttrs
}

type MarkArg interface {
	applyMark(*MarkAttrs, *[]Component)
}

func defaultMarkAttrs() *MarkAttrs {
	return &MarkAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Mark(args ...MarkArg) Node {
	a := defaultMarkAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMark(a, &kids)
	}
	return Node{Tag: "mark", Attrs: a, Kids: kids}
}

func (g Global) applyMark(a *MarkAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyMark(_ *MarkAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyMark(_ *MarkAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *MarkAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
