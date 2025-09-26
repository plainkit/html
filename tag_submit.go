package html

import "strings"

type SubmitAttrs struct {
	Global GlobalAttrs
}

type SubmitArg interface {
	applySubmit(*SubmitAttrs, *[]Component)
}

func defaultSubmitAttrs() *SubmitAttrs {
	return &SubmitAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Submit(args ...SubmitArg) Node {
	a := defaultSubmitAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySubmit(a, &kids)
	}
	return Node{Tag: "submit", Attrs: a, Kids: kids}
}

func (g Global) applySubmit(a *SubmitAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *SubmitAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
