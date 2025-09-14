package blox

import "strings"

type BodyAttrs struct {
	Global GlobalAttrs
}

type BodyArg interface {
	applyBody(*BodyAttrs, *[]Component)
}

func defaultBodyAttrs() *BodyAttrs {
	return &BodyAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Body(args ...BodyArg) Node {
	a := defaultBodyAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBody(a, &kids)
	}
	return Node{Tag: "body", Attrs: a, Kids: kids}
}

// Global option glue
func (g Global) applyBody(a *BodyAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyBody(_ *BodyAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyBody(_ *BodyAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Attrs writer implementation
func (a *BodyAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
}
