package html

import "strings"

type TitleAttrs struct {
	Global GlobalAttrs
}

type TitleArg interface {
	applyTitle(*TitleAttrs, *[]Component)
}

func defaultTitleAttrs() *TitleAttrs {
	return &TitleAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Title(args ...TitleArg) Node {
	a := defaultTitleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTitle(a, &kids)
	}
	return Node{Tag: "title", Attrs: a, Kids: kids}
}

func (g Global) applyTitle(a *TitleAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *TitleAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
