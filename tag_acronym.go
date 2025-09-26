package html

import "strings"

type AcronymAttrs struct {
	Global GlobalAttrs
}

type AcronymArg interface {
	applyAcronym(*AcronymAttrs, *[]Component)
}

func defaultAcronymAttrs() *AcronymAttrs {
	return &AcronymAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Acronym(args ...AcronymArg) Node {
	a := defaultAcronymAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAcronym(a, &kids)
	}
	return Node{Tag: "acronym", Attrs: a, Kids: kids}
}

func (g Global) applyAcronym(a *AcronymAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *AcronymAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
