package html

import "strings"

type DetailsAttrs struct {
	Global GlobalAttrs
	Name   string
	Open   bool
}

type DetailsArg interface {
	applyDetails(*DetailsAttrs, *[]Component)
}

func defaultDetailsAttrs() *DetailsAttrs {
	return &DetailsAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Details(args ...DetailsArg) Node {
	a := defaultDetailsAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDetails(a, &kids)
	}
	return Node{Tag: "details", Attrs: a, Kids: kids}
}

func (g Global) applyDetails(a *DetailsAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyDetails(_ *DetailsAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o NameOpt) applyDetails(a *DetailsAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o OpenOpt) applyDetails(a *DetailsAttrs, _ *[]Component) {
	a.Open = true
}

func (a *DetailsAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Open {
		BoolAttr(sb, "open")
	}
}
