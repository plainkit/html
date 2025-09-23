package html

import "strings"

type SelectedcontentAttrs struct {
	Global GlobalAttrs
}

type SelectedcontentArg interface {
	applySelectedcontent(*SelectedcontentAttrs, *[]Component)
}

func defaultSelectedcontentAttrs() *SelectedcontentAttrs {
	return &SelectedcontentAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Selectedcontent(args ...SelectedcontentArg) Node {
	a := defaultSelectedcontentAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySelectedcontent(a, &kids)
	}
	return Node{Tag: "selectedcontent", Attrs: a, Kids: kids}
}

func (g Global) applySelectedcontent(a *SelectedcontentAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applySelectedcontent(_ *SelectedcontentAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *SelectedcontentAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
