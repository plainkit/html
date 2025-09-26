package html

import "strings"

type H5Attrs struct {
	Global GlobalAttrs
}

type H5Arg interface {
	applyH5(*H5Attrs, *[]Component)
}

func defaultH5Attrs() *H5Attrs {
	return &H5Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H5(args ...H5Arg) Node {
	a := defaultH5Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH5(a, &kids)
	}
	return Node{Tag: "h5", Attrs: a, Kids: kids}
}

func (g Global) applyH5(a *H5Attrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *H5Attrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
