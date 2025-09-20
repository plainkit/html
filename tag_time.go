package html

import "strings"

// Time
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

func (g Global) applyTime(a *TimeAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTime(_ *TimeAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTime(_ *TimeAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o DatetimeOpt) applyTime(a *TimeAttrs, _ *[]Component) { a.Datetime = o.v }

func (a *TimeAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Datetime != "" {
		attr(sb, "datetime", a.Datetime)
	}
}

// Datetime option used by <time>, <del>, and <ins>
type DatetimeOpt struct{ v string }

func Datetime(v string) DatetimeOpt { return DatetimeOpt{v} }
