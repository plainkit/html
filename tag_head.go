package blox

import "strings"

type HeadAttrs struct {
	Global GlobalAttrs
}

type HeadArg interface {
	applyHead(*HeadAttrs, *[]Component)
}

func defaultHeadAttrs() *HeadAttrs {
	return &HeadAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Head(args ...HeadArg) Component {
	a := defaultHeadAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHead(a, &kids)
	}
	return Node{Tag: "head", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyHead(a *HeadAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *HeadAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
