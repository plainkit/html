package html

import "strings"

type MathAttrs struct {
	Global GlobalAttrs
}

type MathArg interface {
	applyMath(*MathAttrs, *[]Component)
}

func defaultMathAttrs() *MathAttrs {
	return &MathAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Math(args ...MathArg) Node {
	a := defaultMathAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMath(a, &kids)
	}
	return Node{Tag: "math", Attrs: a, Kids: kids}
}

func (g Global) applyMath(a *MathAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *MathAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
