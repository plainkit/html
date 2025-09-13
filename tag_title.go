package blox

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
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func HeadTitle(args ...TitleArg) Component {
	a := defaultTitleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTitle(a, &kids)
	}
	return Node{Tag: "title", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyTitle(a *TitleAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyTitle(_ *TitleAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTitle(_ *TitleAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *TitleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
