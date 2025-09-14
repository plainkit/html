package blox

import "strings"

// Nav
type NavAttrs struct {
	Global GlobalAttrs
}

type NavArg interface {
	applyNav(*NavAttrs, *[]Component)
}

func defaultNavAttrs() *NavAttrs {
	return &NavAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Nav(args ...NavArg) Node {
	a := defaultNavAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNav(a, &kids)
	}
	return Node{Tag: "nav", Attrs: a, Kids: kids}
}

func (g Global) applyNav(a *NavAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyNav(_ *NavAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyNav(_ *NavAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *NavAttrs) writeAttrs(sb *strings.Builder)         { writeGlobal(sb, &a.Global) }
