package html

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
			Style:  "",
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

func (g Global) applyBody(a *BodyAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *BodyAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
