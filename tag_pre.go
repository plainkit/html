package html

import "strings"

type PreAttrs struct {
	Global GlobalAttrs
}

type PreArg interface {
	ApplyPre(*PreAttrs, *[]Component)
}

func defaultPreAttrs() *PreAttrs {
	return &PreAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Pre(args ...PreArg) Node {
	a := defaultPreAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyPre(a, &kids)
	}

	return Node{Tag: "pre", Attrs: a, Kids: kids}
}

func (g Global) ApplyPre(a *PreAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *PreAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
