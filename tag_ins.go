package html

import "strings"

// Ins
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

func (g Global) applyIns(a *InsAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyIns(_ *InsAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyIns(_ *InsAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o CiteOpt) applyIns(a *InsAttrs, _ *[]Component)     { a.Cite = o.v }
func (o DatetimeOpt) applyIns(a *InsAttrs, _ *[]Component) { a.Datetime = o.v }

func (a *InsAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cite != "" {
		attr(sb, "cite", a.Cite)
	}
	if a.Datetime != "" {
		attr(sb, "datetime", a.Datetime)
	}
}
