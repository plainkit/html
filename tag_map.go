package html

import "strings"

type MapAttrs struct {
	Global GlobalAttrs
	Name   string
}

type MapArg interface {
	applyMap(*MapAttrs, *[]Component)
}

func defaultMapAttrs() *MapAttrs {
	return &MapAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Map(args ...MapArg) Node {
	a := defaultMapAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMap(a, &kids)
	}
	return Node{Tag: "map", Attrs: a, Kids: kids}
}

func (g Global) applyMap(a *MapAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyMap(_ *MapAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyMap(_ *MapAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o NameOpt) applyMap(a *MapAttrs, _ *[]Component) {
	a.Name = o.v
}

func (a *MapAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
}
