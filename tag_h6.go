package html

import "strings"

type H6Attrs struct {
	Global GlobalAttrs
}

type H6Arg interface {
	ApplyH6(*H6Attrs, *[]Component)
}

func defaultH6Attrs() *H6Attrs {
	return &H6Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H6(args ...H6Arg) Node {
	a := defaultH6Attrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyH6(a, &kids)
	}

	return Node{Tag: "h6", Attrs: a, Kids: kids}
}

func (g Global) ApplyH6(a *H6Attrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *H6Attrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
