package html

import "strings"

type PreAttrs struct {
	Global GlobalAttrs
	Width  string
}

type PreArg interface {
	applyPre(*PreAttrs, *[]Component)
}

func defaultPreAttrs() *PreAttrs {
	return &PreAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Pre(args ...PreArg) Node {
	a := defaultPreAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPre(a, &kids)
	}
	return Node{Tag: "pre", Attrs: a, Kids: kids}
}

func (g Global) applyPre(a *PreAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o WidthOpt) applyPre(a *PreAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *PreAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
