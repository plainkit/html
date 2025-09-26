package html

import "strings"

type RbAttrs struct {
	Global GlobalAttrs
}

type RbArg interface {
	applyRb(*RbAttrs, *[]Component)
}

func defaultRbAttrs() *RbAttrs {
	return &RbAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rb(args ...RbArg) Node {
	a := defaultRbAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRb(a, &kids)
	}
	return Node{Tag: "rb", Attrs: a, Kids: kids}
}

func (g Global) applyRb(a *RbAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *RbAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
