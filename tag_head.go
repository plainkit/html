package html

import "strings"

type HeadAttrs struct {
	Global GlobalAttrs
}

type HeadArg interface {
	ApplyHead(*HeadAttrs, *[]Component)
}

func defaultHeadAttrs() *HeadAttrs {
	return &HeadAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Head(args ...HeadArg) Node {
	a := defaultHeadAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyHead(a, &kids)
	}
	return Node{Tag: "head", Attrs: a, Kids: kids}
}

func (g Global) ApplyHead(a *HeadAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *HeadAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
