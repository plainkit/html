package html

import "strings"

type CenterAttrs struct {
	Global GlobalAttrs
}

type CenterArg interface {
	applyCenter(*CenterAttrs, *[]Component)
}

func defaultCenterAttrs() *CenterAttrs {
	return &CenterAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Center(args ...CenterArg) Node {
	a := defaultCenterAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyCenter(a, &kids)
	}
	return Node{Tag: "center", Attrs: a, Kids: kids}
}

func (g Global) applyCenter(a *CenterAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *CenterAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
