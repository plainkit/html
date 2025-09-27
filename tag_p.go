package html

import "strings"

type PAttrs struct {
	Global GlobalAttrs
	Align  string
}

type PArg interface {
	applyP(*PAttrs, *[]Component)
}

func defaultPAttrs() *PAttrs {
	return &PAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func P(args ...PArg) Node {
	a := defaultPAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyP(a, &kids)
	}
	return Node{Tag: "p", Attrs: a, Kids: kids}
}

func (g Global) applyP(a *PAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AlignOpt) applyP(a *PAttrs, _ *[]Component) {
	a.Align = o.v
}

func (a *PAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
}
