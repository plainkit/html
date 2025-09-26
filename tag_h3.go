package html

import "strings"

type H3Attrs struct {
	Global GlobalAttrs
	Align  string
}

type H3Arg interface {
	applyH3(*H3Attrs, *[]Component)
}

func defaultH3Attrs() *H3Attrs {
	return &H3Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H3(args ...H3Arg) Node {
	a := defaultH3Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH3(a, &kids)
	}
	return Node{Tag: "h3", Attrs: a, Kids: kids}
}

func (g Global) applyH3(a *H3Attrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyH3(a *H3Attrs, _ *[]Component) {
	a.Align = o.v
}

func (a *H3Attrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
