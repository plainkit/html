package html

import "strings"

type NavAttrs struct {
	Global GlobalAttrs
}

type NavArg interface {
	ApplyNav(*NavAttrs, *[]Component)
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
		ar.ApplyNav(a, &kids)
	}
	return Node{Tag: "nav", Attrs: a, Kids: kids}
}

func (g Global) ApplyNav(a *NavAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *NavAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
