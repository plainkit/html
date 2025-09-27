package html

import "strings"

type PreAttrs struct {
	Global GlobalAttrs
}

type PreArg interface {
	applyPre(*PreAttrs, *[]Component)
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
		ar.applyPre(a, &kids)
	}
	return Node{Tag: "pre", Attrs: a, Kids: kids}
}

func (g Global) applyPre(a *PreAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *PreAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
