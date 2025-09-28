package html

import "strings"

type RpAttrs struct {
	Global GlobalAttrs
}

type RpArg interface {
	ApplyRp(*RpAttrs, *[]Component)
}

func defaultRpAttrs() *RpAttrs {
	return &RpAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rp(args ...RpArg) Node {
	a := defaultRpAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyRp(a, &kids)
	}
	return Node{Tag: "rp", Attrs: a, Kids: kids}
}

func (g Global) ApplyRp(a *RpAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *RpAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
