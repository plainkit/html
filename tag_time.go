package html

import "strings"

type TimeAttrs struct {
	Global   GlobalAttrs
	Datetime string
}

type TimeArg interface {
	applyTime(*TimeAttrs, *[]Component)
}

func defaultTimeAttrs() *TimeAttrs {
	return &TimeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Time(args ...TimeArg) Node {
	a := defaultTimeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTime(a, &kids)
	}
	return Node{Tag: "time", Attrs: a, Kids: kids}
}

func (g Global) applyTime(a *TimeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DatetimeOpt) applyTime(a *TimeAttrs, _ *[]Component) {
	a.Datetime = o.v
}

func (a *TimeAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Datetime != "" {
		Attr(sb, "datetime", a.Datetime)
	}
}
