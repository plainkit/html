package blox

import "strings"

// CitationTag (avoiding conflict with Cite option)
type CitationAttrs struct {
	Global GlobalAttrs
}

type CitationArg interface {
	applyCitation(*CitationAttrs, *[]Component)
}

func defaultCitationAttrs() *CitationAttrs {
	return &CitationAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Citation(args ...CitationArg) Component {
	a := defaultCitationAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCitation(a, &kids)
	}
	return Node{Tag: "cite", Attrs: a, Kids: kids}
}

func (g Global) applyCitation(a *CitationAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyCitation(_ *CitationAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyCitation(_ *CitationAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *CitationAttrs) writeAttrs(sb *strings.Builder)              { writeGlobal(sb, &a.Global) }
