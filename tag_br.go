package html

import "strings"

type BrAttrs struct {
	Global GlobalAttrs
}

type BrArg interface {
	ApplyBr(*BrAttrs, *[]Component)
}

func defaultBrAttrs() *BrAttrs {
	return &BrAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Br(args ...BrArg) Node {
	a := defaultBrAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyBr(a, &kids)
	}

	return Node{Tag: "br", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyBr(a *BrAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *BrAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
