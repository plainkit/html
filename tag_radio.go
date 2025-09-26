package html

import "strings"

type RadioAttrs struct {
	Global GlobalAttrs
}

type RadioArg interface {
	applyRadio(*RadioAttrs, *[]Component)
}

func defaultRadioAttrs() *RadioAttrs {
	return &RadioAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Radio(args ...RadioArg) Node {
	a := defaultRadioAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRadio(a, &kids)
	}
	return Node{Tag: "radio", Attrs: a, Kids: kids}
}

func (g Global) applyRadio(a *RadioAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *RadioAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
