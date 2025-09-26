package html

import "strings"

type NoframesAttrs struct {
	Global GlobalAttrs
}

type NoframesArg interface {
	applyNoframes(*NoframesAttrs, *[]Component)
}

func defaultNoframesAttrs() *NoframesAttrs {
	return &NoframesAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Noframes(args ...NoframesArg) Node {
	a := defaultNoframesAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyNoframes(a, &kids)
	}
	return Node{Tag: "noframes", Attrs: a, Kids: kids}
}

func (g Global) applyNoframes(a *NoframesAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *NoframesAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
