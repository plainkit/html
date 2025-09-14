package blox

import "strings"

// Legend
type LegendAttrs struct {
	Global GlobalAttrs
}

type LegendArg interface {
	applyLegend(*LegendAttrs, *[]Component)
}

func defaultLegendAttrs() *LegendAttrs {
	return &LegendAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
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
		ar.applyLegend(a, &kids)
	}
	return Node{Tag: "legend", Attrs: a, Kids: kids}
}

func (g Global) applyLegend(a *LegendAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyLegend(_ *LegendAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyLegend(_ *LegendAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *LegendAttrs) writeAttrs(sb *strings.Builder)            { writeGlobal(sb, &a.Global) }
