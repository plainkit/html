package html

import "strings"

// Summary
type SummaryAttrs struct {
	Global GlobalAttrs
}

type SummaryArg interface {
	applySummary(*SummaryAttrs, *[]Component)
}

func defaultSummaryAttrs() *SummaryAttrs {
	return &SummaryAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Summary(args ...SummaryArg) Node {
	a := defaultSummaryAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySummary(a, &kids)
	}
	return Node{Tag: "summary", Attrs: a, Kids: kids}
}

func (g Global) applySummary(a *SummaryAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applySummary(_ *SummaryAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applySummary(_ *SummaryAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *SummaryAttrs) writeAttrs(sb *strings.Builder)             { writeGlobal(sb, &a.Global) }
