package html

import "strings"

type MonthAttrs struct {
	Global GlobalAttrs
}

type MonthArg interface {
	applyMonth(*MonthAttrs, *[]Component)
}

func defaultMonthAttrs() *MonthAttrs {
	return &MonthAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Month(args ...MonthArg) Node {
	a := defaultMonthAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMonth(a, &kids)
	}
	return Node{Tag: "month", Attrs: a, Kids: kids}
}

func (g Global) applyMonth(a *MonthAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *MonthAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
