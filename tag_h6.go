package html

import "strings"

type H6Attrs struct {
	Global GlobalAttrs
	Align  string
}

type H6Arg interface {
	applyH6(*H6Attrs, *[]Component)
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
		ar.applyH6(a, &kids)
	}
	return Node{Tag: "h6", Attrs: a, Kids: kids}
}

func (g Global) applyH6(a *H6Attrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyH6(a *H6Attrs, _ *[]Component) {
	a.Align = o.v
}

func (a *H6Attrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
