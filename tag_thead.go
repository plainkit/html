package html

import "strings"

// Thead
type TheadAttrs struct {
	Global GlobalAttrs
}

type TheadArg interface {
	applyThead(*TheadAttrs, *[]Component)
}

func defaultTheadAttrs() *TheadAttrs {
	return &TheadAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Thead(args ...TheadArg) Node {
	a := defaultTheadAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyThead(a, &kids)
	}
	return Node{Tag: "thead", Attrs: a, Kids: kids}
}

func (g Global) applyThead(a *TheadAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyThead(_ *TheadAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyThead(_ *TheadAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TheadAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
