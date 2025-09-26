package html

import "strings"

type H2Attrs struct {
	Global GlobalAttrs
}

type H2Arg interface {
	applyH2(*H2Attrs, *[]Component)
}

func defaultH2Attrs() *H2Attrs {
	return &H2Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H2(args ...H2Arg) Node {
	a := defaultH2Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH2(a, &kids)
	}
	return Node{Tag: "h2", Attrs: a, Kids: kids}
}

func (g Global) applyH2(a *H2Attrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *H2Attrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
