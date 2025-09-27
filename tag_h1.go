package html

import "strings"

type H1Attrs struct {
	Global GlobalAttrs
	Align  string
}

type H1Arg interface {
	applyH1(*H1Attrs, *[]Component)
}

func defaultH1Attrs() *H1Attrs {
	return &H1Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H1(args ...H1Arg) Node {
	a := defaultH1Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH1(a, &kids)
	}
	return Node{Tag: "h1", Attrs: a, Kids: kids}
}

func (g Global) applyH1(a *H1Attrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AlignOpt) applyH1(a *H1Attrs, _ *[]Component) {
	a.Align = o.v
}

func (a *H1Attrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
