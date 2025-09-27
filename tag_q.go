package html

import "strings"

type QAttrs struct {
	Global GlobalAttrs
	Cite   string
}

type QArg interface {
	applyQ(*QAttrs, *[]Component)
}

func defaultQAttrs() *QAttrs {
	return &QAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Q(args ...QArg) Node {
	a := defaultQAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyQ(a, &kids)
	}
	return Node{Tag: "q", Attrs: a, Kids: kids}
}

func (g Global) applyQ(a *QAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CiteOpt) applyQ(a *QAttrs, _ *[]Component) {
	a.Cite = o.v
}

func (a *QAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Cite != "" {
		Attr(sb, "cite", a.Cite)
	}
}
