package html

import "strings"

type InsAttrs struct {
	Global   GlobalAttrs
	Cite     string
	Datetime string
}

type InsArg interface {
	applyIns(*InsAttrs, *[]Component)
}

func defaultInsAttrs() *InsAttrs {
	return &InsAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ins(args ...InsArg) Node {
	a := defaultInsAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyIns(a, &kids)
	}
	return Node{Tag: "ins", Attrs: a, Kids: kids}
}

func (g Global) applyIns(a *InsAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CiteOpt) applyIns(a *InsAttrs, _ *[]Component) {
	a.Cite = o.v
}
func (o DatetimeOpt) applyIns(a *InsAttrs, _ *[]Component) {
	a.Datetime = o.v
}

func (a *InsAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Cite != "" {
		Attr(sb, "cite", a.Cite)
	}
	if a.Datetime != "" {
		Attr(sb, "datetime", a.Datetime)
	}
}
