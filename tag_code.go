package html

import "strings"

type CodeAttrs struct {
	Global GlobalAttrs
}

type CodeArg interface {
	applyCode(*CodeAttrs, *[]Component)
}

func defaultCodeAttrs() *CodeAttrs {
	return &CodeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Code(args ...CodeArg) Node {
	a := defaultCodeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCode(a, &kids)
	}
	return Node{Tag: "code", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyCode(a *CodeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyCode(_ *CodeAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyCode(_ *CodeAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *CodeAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
