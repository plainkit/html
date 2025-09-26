package html

import "strings"

type WeekAttrs struct {
	Global GlobalAttrs
}

type WeekArg interface {
	applyWeek(*WeekAttrs, *[]Component)
}

func defaultWeekAttrs() *WeekAttrs {
	return &WeekAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Week(args ...WeekArg) Node {
	a := defaultWeekAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyWeek(a, &kids)
	}
	return Node{Tag: "week", Attrs: a, Kids: kids}
}

func (g Global) applyWeek(a *WeekAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *WeekAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
