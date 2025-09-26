package html

import "strings"

type EmAttrs struct {
	Global GlobalAttrs
}

type EmArg interface {
	applyEm(*EmAttrs, *[]Component)
}

func defaultEmAttrs() *EmAttrs {
	return &EmAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Em(args ...EmArg) Node {
	a := defaultEmAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyEm(a, &kids)
	}
	return Node{Tag: "em", Attrs: a, Kids: kids}
}

func (g Global) applyEm(a *EmAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *EmAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
