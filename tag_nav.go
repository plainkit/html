package html

import "strings"

type NavAttrs struct {
	Global GlobalAttrs
}

type NavArg interface {
	applyNav(*NavAttrs, *[]Component)
}

func defaultNavAttrs() *NavAttrs {
	return &NavAttrs{
		Global: GlobalAttrs{
			Style:  "",
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

func (g Global) applyNav(a *NavAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NavAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
