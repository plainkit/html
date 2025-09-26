package html

import "strings"

type PlaintextAttrs struct {
	Global GlobalAttrs
}

type PlaintextArg interface {
	applyPlaintext(*PlaintextAttrs, *[]Component)
}

func defaultPlaintextAttrs() *PlaintextAttrs {
	return &PlaintextAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Plaintext(args ...PlaintextArg) Node {
	a := defaultPlaintextAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPlaintext(a, &kids)
	}
	return Node{Tag: "plaintext", Attrs: a, Kids: kids}
}

func (g Global) applyPlaintext(a *PlaintextAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *PlaintextAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
