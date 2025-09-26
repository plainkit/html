package html

import "strings"

type SummaryAttrs struct {
	Global          GlobalAttrs
	DisplayListItem string
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

func (g Global) applySummary(a *SummaryAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DisplayListItemOpt) applySummary(a *SummaryAttrs, _ *[]Component) {
	a.DisplayListItem = o.v
}

func (a *SummaryAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.DisplayListItem != "" {
		Attr(sb, "display_list_item", a.DisplayListItem)
	}
}
