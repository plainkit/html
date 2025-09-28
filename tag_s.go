package html

import "strings"

type SAttrs struct {
	Global GlobalAttrs
}

type SArg interface {
	ApplyS(*SAttrs, *[]Component)
}

func defaultSAttrs() *SAttrs {
	return &SAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func S(args ...SArg) Node {
	a := defaultSAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyS(a, &kids)
	}
	return Node{Tag: "s", Attrs: a, Kids: kids}
}

func (g Global) ApplyS(a *SAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
