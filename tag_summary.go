package html

import "strings"

type SummaryAttrs struct {
	Global GlobalAttrs
}

type SummaryArg interface {
	ApplySummary(*SummaryAttrs, *[]Component)
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
		ar.ApplySummary(a, &kids)
	}
	return Node{Tag: "summary", Attrs: a, Kids: kids}
}

func (g Global) ApplySummary(a *SummaryAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SummaryAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
