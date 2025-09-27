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

func (a *WbrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
