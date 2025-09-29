package html

import "strings"

type CiteAttrs struct {
	Global GlobalAttrs
}

type CiteArg interface {
	ApplyCite(*CiteAttrs, *[]Component)
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
		ar.ApplyCite(a, &kids)
	}

	return Node{Tag: "cite", Attrs: a, Kids: kids}
}

func (g Global) ApplyCite(a *CiteAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *CiteAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
