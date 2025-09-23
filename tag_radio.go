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

func (o TxtOpt) applyRadio(_ *RadioAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyRadio(_ *RadioAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *RadioAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
