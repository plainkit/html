package html

import "strings"

// Section
type SectionAttrs struct {
	Global GlobalAttrs
}

type SectionArg interface {
	applySection(*SectionAttrs, *[]Component)
}

func defaultSectionAttrs() *SectionAttrs {
	return &SectionAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Section(args ...SectionArg) Node {
	a := defaultSectionAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applySection(a, &kids)
	}
	return Node{Tag: "section", Attrs: a, Kids: kids}
}

func (g Global) applySection(a *SectionAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applySection(_ *SectionAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applySection(_ *SectionAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *SectionAttrs) writeAttrs(sb *strings.Builder)             { writeGlobal(sb, &a.Global) }
