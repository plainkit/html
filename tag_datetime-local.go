package html

import "strings"

type DatetimeLocalAttrs struct {
	Global GlobalAttrs
}

type DatetimeLocalArg interface {
	applyDatetimeLocal(*DatetimeLocalAttrs, *[]Component)
}

func defaultDatetimeLocalAttrs() *DatetimeLocalAttrs {
	return &DatetimeLocalAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func DatetimeLocal(args ...DatetimeLocalArg) Node {
	a := defaultDatetimeLocalAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDatetimeLocal(a, &kids)
	}
	return Node{Tag: "datetime-local", Attrs: a, Kids: kids}
}

func (g Global) applyDatetimeLocal(a *DatetimeLocalAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyDatetimeLocal(_ *DatetimeLocalAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyDatetimeLocal(_ *DatetimeLocalAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *DatetimeLocalAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
