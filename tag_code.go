package html

import "strings"

type CodeAttrs struct {
	Global GlobalAttrs
}

type CodeArg interface {
	ApplyCode(*CodeAttrs, *[]Component)
}

func defaultCodeAttrs() *CodeAttrs {
	return &CodeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Code(args ...CodeArg) Node {
	a := defaultCodeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyCode(a, &kids)
	}
	return Node{Tag: "code", Attrs: a, Kids: kids}
}

func (g Global) ApplyCode(a *CodeAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *CodeAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
