package html

import "strings"

type StrongAttrs struct {
	Global GlobalAttrs
}

type StrongArg interface {
	ApplyStrong(*StrongAttrs, *[]Component)
}

func defaultStrongAttrs() *StrongAttrs {
	return &StrongAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Strong(args ...StrongArg) Node {
	a := defaultStrongAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyStrong(a, &kids)
	}
	return Node{Tag: "strong", Attrs: a, Kids: kids}
}

func (g Global) ApplyStrong(a *StrongAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *StrongAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
