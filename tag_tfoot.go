package html

import "strings"

type TfootAttrs struct {
	Global GlobalAttrs
}

type TfootArg interface {
	applyTfoot(*TfootAttrs, *[]Component)
}

func defaultTfootAttrs() *TfootAttrs {
	return &TfootAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tfoot(args ...TfootArg) Node {
	a := defaultTfootAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTfoot(a, &kids)
	}
	return Node{Tag: "tfoot", Attrs: a, Kids: kids}
}

func (g Global) applyTfoot(a *TfootAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *TfootAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
