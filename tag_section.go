package html

import "strings"

type SectionAttrs struct {
	Global GlobalAttrs
}

type SectionArg interface {
	ApplySection(*SectionAttrs, *[]Component)
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
		ar.ApplySection(a, &kids)
	}

	return Node{Tag: "section", Attrs: a, Kids: kids}
}

func (g Global) ApplySection(a *SectionAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SectionAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
