package html

import "strings"

type H4Attrs struct {
	Global GlobalAttrs
}

type H4Arg interface {
	applyH4(*H4Attrs, *[]Component)
}

func defaultH4Attrs() *H4Attrs {
	return &H4Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H4(args ...H4Arg) Node {
	a := defaultH4Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH4(a, &kids)
	}
	return Node{Tag: "h4", Attrs: a, Kids: kids}
}

func (g Global) applyH4(a *H4Attrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *H4Attrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
