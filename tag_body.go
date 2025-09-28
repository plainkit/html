package html

import "strings"

type BodyAttrs struct {
	Global GlobalAttrs
}

type BodyArg interface {
	ApplyBody(*BodyAttrs, *[]Component)
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
		ar.ApplyBody(a, &kids)
	}
	return Node{Tag: "body", Attrs: a, Kids: kids}
}

func (g Global) ApplyBody(a *BodyAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *BodyAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
