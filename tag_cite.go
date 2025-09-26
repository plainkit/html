package html

import "strings"

type CiteAttrs struct {
	Global GlobalAttrs
}

type CiteArg interface {
	applyCite(*CiteAttrs, *[]Component)
}

func defaultCiteAttrs() *CiteAttrs {
	return &CiteAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Cite(args ...CiteArg) Node {
	a := defaultCiteAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCite(a, &kids)
	}
	return Node{Tag: "cite", Attrs: a, Kids: kids}
}

func (g Global) applyCite(a *CiteAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *CiteAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
