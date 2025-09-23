package html

import "strings"

type WbrAttrs struct {
	Global GlobalAttrs
}

type WbrArg interface {
	applyWbr(*WbrAttrs, *[]Component)
}

func defaultWbrAttrs() *WbrAttrs {
	return &WbrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Wbr(args ...WbrArg) Node {
	a := defaultWbrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyWbr(a, &kids)
	}
	return Node{Tag: "wbr", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyWbr(a *WbrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyWbr(_ *WbrAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyWbr(_ *WbrAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *WbrAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
