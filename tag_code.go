package html

import "strings"

type CodeAttrs struct {
	Global GlobalAttrs
}

type CodeArg interface {
	applyCode(*CodeAttrs, *[]Component)
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
		ar.applyCode(a, &kids)
	}
	return Node{Tag: "code", Attrs: a, Kids: kids}
}

func (g Global) applyCode(a *CodeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *CodeAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
