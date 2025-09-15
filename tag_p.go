package html

import "strings"

type PAttrs struct {
	Global GlobalAttrs
}

type PArg interface {
	applyP(*PAttrs, *[]Component)
}

func defaultPAttrs() *PAttrs {
	return &PAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func P(args ...PArg) Node {
	a := defaultPAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyP(a, &kids)
	}
	return Node{Tag: "p", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyP(a *PAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyP(_ *PAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyP(_ *PAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *PAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
