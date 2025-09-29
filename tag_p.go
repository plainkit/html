package html

import "strings"

type PAttrs struct {
	Global GlobalAttrs
}

type PArg interface {
	ApplyP(*PAttrs, *[]Component)
}

func defaultPAttrs() *PAttrs {
	return &PAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyP(a, &kids)
	}

	return Node{Tag: "p", Attrs: a, Kids: kids}
}

func (g Global) ApplyP(a *PAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *PAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
