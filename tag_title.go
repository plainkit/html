package html

import "strings"

type TitleAttrs struct {
	Global GlobalAttrs
}

type TitleArg interface {
	ApplyTitle(*TitleAttrs, *[]Component)
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
		ar.ApplyTitle(a, &kids)
	}
	return Node{Tag: "title", Attrs: a, Kids: kids}
}

func (g Global) ApplyTitle(a *TitleAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *TitleAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
