package html

import "strings"

type EmAttrs struct {
	Global GlobalAttrs
}

type EmArg interface {
	ApplyEm(*EmAttrs, *[]Component)
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
		ar.ApplyEm(a, &kids)
	}
	return Node{Tag: "em", Attrs: a, Kids: kids}
}

func (g Global) ApplyEm(a *EmAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *EmAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
