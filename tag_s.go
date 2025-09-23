package html

import "strings"

type SAttrs struct {
	Global GlobalAttrs
}

type SArg interface {
	applyS(*SAttrs, *[]Component)
}

func defaultSAttrs() *SAttrs {
	return &SAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func S(args ...SArg) Node {
	a := defaultSAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyS(a, &kids)
	}
	return Node{Tag: "s", Attrs: a, Kids: kids}
}

func (g Global) applyS(a *SAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyS(_ *SAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyS(_ *SAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *SAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
