package html

import "strings"

type DatalistAttrs struct {
	Global GlobalAttrs
}

type DatalistArg interface {
	ApplyDatalist(*DatalistAttrs, *[]Component)
}

func defaultDatalistAttrs() *DatalistAttrs {
	return &DatalistAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Datalist(args ...DatalistArg) Node {
	a := defaultDatalistAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyDatalist(a, &kids)
	}
	return Node{Tag: "datalist", Attrs: a, Kids: kids}
}

func (g Global) ApplyDatalist(a *DatalistAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *DatalistAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
