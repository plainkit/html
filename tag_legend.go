package html

import "strings"

type LegendAttrs struct {
	Global GlobalAttrs
}

type LegendArg interface {
	ApplyLegend(*LegendAttrs, *[]Component)
}

func defaultLegendAttrs() *LegendAttrs {
	return &LegendAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Legend(args ...LegendArg) Node {
	a := defaultLegendAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyLegend(a, &kids)
	}
	return Node{Tag: "legend", Attrs: a, Kids: kids}
}

func (g Global) ApplyLegend(a *LegendAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *LegendAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
