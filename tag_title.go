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

type TitleComponent Node

func (title TitleComponent) render(sb *strings.Builder) {
	Node(title).render(sb)
}

func HeadTitle(args ...TitleArg) TitleComponent {
	a := defaultTitleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTitle(a, &kids)
	}
	return TitleComponent{Tag: "title", Attrs: a, Kids: kids}
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

// Compile-time type safety: Title can be added to Head
func (title TitleComponent) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, title)
}

// Attrs writer implementation
func (a *TitleAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
