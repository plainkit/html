package html

import "strings"

type DataAttrs struct {
	Global GlobalAttrs
	Value  string
}

type DataArg interface {
	applyData(*DataAttrs, *[]Component)
}

func defaultDataAttrs() *DataAttrs {
	return &DataAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Data(args ...DataArg) Node {
	a := defaultDataAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyData(a, &kids)
	}
	return Node{Tag: "data", Attrs: a, Kids: kids}
}

func (g Global) applyData(a *DataAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyData(_ *DataAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyData(_ *DataAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o ValueOpt) applyData(a *DataAttrs, _ *[]Component) {
	a.Value = o.v
}

func (a *DataAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Value != "" {
		Attr(sb, "value", a.Value)
	}
}
