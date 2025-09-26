package html

import "strings"

type NoembedAttrs struct {
	Global GlobalAttrs
}

type NoembedArg interface {
	applyNoembed(*NoembedAttrs, *[]Component)
}

func defaultNoembedAttrs() *NoembedAttrs {
	return &NoembedAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Noembed(args ...NoembedArg) Node {
	a := defaultNoembedAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNoembed(a, &kids)
	}
	return Node{Tag: "noembed", Attrs: a, Kids: kids}
}

func (g Global) applyNoembed(a *NoembedAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NoembedAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
