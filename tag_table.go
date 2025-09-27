package html

import "strings"

type TableAttrs struct {
	Global GlobalAttrs
	Border string
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
	g.do(&a.Global)
}

func (o BorderOpt) applyTable(a *TableAttrs, _ *[]Component) {
	a.Border = o.v
}

func (a *TableAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Border != "" {
		Attr(sb, "border", a.Border)
	}
}
