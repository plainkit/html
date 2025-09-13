package blox

import "strings"

type PreAttrs struct {
	Global GlobalAttrs
}

type PreArg interface {
	applyPre(*PreAttrs, *[]Component)
}

func defaultPreAttrs() *PreAttrs {
	return &PreAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Pre(args ...PreArg) Component {
	a := defaultPreAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPre(a, &kids)
	}
	return Node{Tag: "pre", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyPre(a *PreAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyPre(_ *PreAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyPre(_ *PreAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *PreAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
