package html

import "strings"

type HiddenAttrs struct {
	Global GlobalAttrs
}

type HiddenArg interface {
	applyHidden(*HiddenAttrs, *[]Component)
}

func defaultHiddenAttrs() *HiddenAttrs {
	return &HiddenAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Hidden(args ...HiddenArg) Node {
	a := defaultHiddenAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHidden(a, &kids)
	}
	return Node{Tag: "hidden", Attrs: a, Kids: kids}
}

func (g Global) applyHidden(a *HiddenAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyHidden(_ *HiddenAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyHidden(_ *HiddenAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *HiddenAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
