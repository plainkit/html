package html

import "strings"

type TableAttrs struct {
	Global GlobalAttrs
}

type TableArg interface {
	applyTable(*TableAttrs, *[]Component)
}

func defaultTableAttrs() *TableAttrs {
	return &TableAttrs{
		Global: GlobalAttrs{
			Style:  "",
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

func (g Global) applyTable(a *TableAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *TableAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
