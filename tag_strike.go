package html

import "strings"

type StrikeAttrs struct {
	Global GlobalAttrs
}

type StrikeArg interface {
	applyStrike(*StrikeAttrs, *[]Component)
}

func defaultStrikeAttrs() *StrikeAttrs {
	return &StrikeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Strike(args ...StrikeArg) Node {
	a := defaultStrikeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyStrike(a, &kids)
	}
	return Node{Tag: "strike", Attrs: a, Kids: kids}
}

func (g Global) applyStrike(a *StrikeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyStrike(_ *StrikeAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyStrike(_ *StrikeAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *StrikeAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
