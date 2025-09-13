package blox

import "strings"

// Datalist
type DatalistAttrs struct {
	Global GlobalAttrs
}

type DatalistArg interface {
	applyDatalist(*DatalistAttrs, *[]Component)
}

func defaultDatalistAttrs() *DatalistAttrs {
	return &DatalistAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Datalist(args ...DatalistArg) Component {
	a := defaultDatalistAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDatalist(a, &kids)
	}
	return Node{Tag: "datalist", Attrs: a, Kids: kids}
}

func (g Global) applyDatalist(a *DatalistAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyDatalist(_ *DatalistAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyDatalist(_ *DatalistAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *DatalistAttrs) writeAttrs(sb *strings.Builder)              { writeGlobal(sb, &a.Global) }
