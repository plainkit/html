package html

import "strings"

type VarAttrs struct {
	Global GlobalAttrs
}

type VarArg interface {
	ApplyVar(*VarAttrs, *[]Component)
}

func defaultVarAttrs() *VarAttrs {
	return &VarAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Var(args ...VarArg) Node {
	a := defaultVarAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyVar(a, &kids)
	}

	return Node{Tag: "var", Attrs: a, Kids: kids}
}

func (g Global) ApplyVar(a *VarAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *VarAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
