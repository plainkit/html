package blox

import "strings"

// Var
type VarAttrs struct {
	Global GlobalAttrs
}

type VarArg interface {
	applyVar(*VarAttrs, *[]Component)
}

func defaultVarAttrs() *VarAttrs {
	return &VarAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
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
		ar.applyVar(a, &kids)
	}
	return Node{Tag: "var", Attrs: a, Kids: kids}
}

func (g Global) applyVar(a *VarAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyVar(_ *VarAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyVar(_ *VarAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *VarAttrs) writeAttrs(sb *strings.Builder)         { writeGlobal(sb, &a.Global) }
