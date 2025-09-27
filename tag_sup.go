package html

import "strings"

type SupAttrs struct {
	Global GlobalAttrs
}

type SupArg interface {
	applySup(*SupAttrs, *[]Component)
}

func defaultSupAttrs() *SupAttrs {
	return &SupAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Sup(args ...SupArg) Node {
	a := defaultSupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySup(a, &kids)
	}
	return Node{Tag: "sup", Attrs: a, Kids: kids}
}

func (g Global) applySup(a *SupAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SupAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
