package html

import "strings"

type HeaderAttrs struct {
	Global GlobalAttrs
}

type HeaderArg interface {
	applyHeader(*HeaderAttrs, *[]Component)
}

func defaultHeaderAttrs() *HeaderAttrs {
	return &HeaderAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Header(args ...HeaderArg) Node {
	a := defaultHeaderAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHeader(a, &kids)
	}
	return Node{Tag: "header", Attrs: a, Kids: kids}
}

func (g Global) applyHeader(a *HeaderAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *HeaderAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
