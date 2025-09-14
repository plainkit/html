package blox

import "strings"

// Table
type TableAttrs struct {
	Global GlobalAttrs
}

type TableArg interface {
	applyTable(*TableAttrs, *[]Component)
}

func defaultTableAttrs() *TableAttrs {
	return &TableAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Table(args ...TableArg) Node {
	a := defaultTableAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTable(a, &kids)
	}
	return Node{Tag: "table", Attrs: a, Kids: kids}
}

func (g Global) applyTable(a *TableAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTable(_ *TableAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTable(_ *TableAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TableAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
