package html

import "strings"

type DateAttrs struct {
	Global GlobalAttrs
}

type DateArg interface {
	applyDate(*DateAttrs, *[]Component)
}

func defaultDateAttrs() *DateAttrs {
	return &DateAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Date(args ...DateArg) Node {
	a := defaultDateAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDate(a, &kids)
	}
	return Node{Tag: "date", Attrs: a, Kids: kids}
}

func (g Global) applyDate(a *DateAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *DateAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
