package html

import "strings"

type SmallAttrs struct {
	Global GlobalAttrs
}

type SmallArg interface {
	applySmall(*SmallAttrs, *[]Component)
}

func defaultSmallAttrs() *SmallAttrs {
	return &SmallAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Small(args ...SmallArg) Node {
	a := defaultSmallAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySmall(a, &kids)
	}
	return Node{Tag: "small", Attrs: a, Kids: kids}
}

func (g Global) applySmall(a *SmallAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *SmallAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
