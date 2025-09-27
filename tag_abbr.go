package html

import "strings"

type AbbrAttrs struct {
	Global GlobalAttrs
}

type AbbrArg interface {
	applyAbbr(*AbbrAttrs, *[]Component)
}

func defaultAbbrAttrs() *AbbrAttrs {
	return &AbbrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Abbr(args ...AbbrArg) Node {
	a := defaultAbbrAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAbbr(a, &kids)
	}
	return Node{Tag: "abbr", Attrs: a, Kids: kids}
}

func (g Global) applyAbbr(a *AbbrAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *AbbrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
