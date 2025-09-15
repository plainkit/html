package html

import "strings"

// Del
type DelAttrs struct {
	Global   GlobalAttrs
	Cite     string
	Datetime string
}

type DelArg interface {
	applyDel(*DelAttrs, *[]Component)
}

func defaultDelAttrs() *DelAttrs {
	return &DelAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Del(args ...DelArg) Node {
	a := defaultDelAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDel(a, &kids)
	}
	return Node{Tag: "del", Attrs: a, Kids: kids}
}

func (g Global) applyDel(a *DelAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyDel(_ *DelAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyDel(_ *DelAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (o CiteOpt) applyDel(a *DelAttrs, _ *[]Component)     { a.Cite = o.v }
func (o DatetimeOpt) applyDel(a *DelAttrs, _ *[]Component) { a.Datetime = o.v }

func (a *DelAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cite != "" {
		attr(sb, "cite", a.Cite)
	}
	if a.Datetime != "" {
		attr(sb, "datetime", a.Datetime)
	}
}
