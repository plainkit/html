package html

import "strings"

type ResetAttrs struct {
	Global GlobalAttrs
}

type ResetArg interface {
	applyReset(*ResetAttrs, *[]Component)
}

func defaultResetAttrs() *ResetAttrs {
	return &ResetAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Reset(args ...ResetArg) Node {
	a := defaultResetAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyReset(a, &kids)
	}
	return Node{Tag: "reset", Attrs: a, Kids: kids}
}

func (g Global) applyReset(a *ResetAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *ResetAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
