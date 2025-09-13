package blox

import "strings"

// Tfoot
type TfootAttrs struct {
	Global GlobalAttrs
}

type TfootArg interface {
	applyTfoot(*TfootAttrs, *[]Component)
}

func defaultTfootAttrs() *TfootAttrs {
	return &TfootAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tfoot(args ...TfootArg) Component {
	a := defaultTfootAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTfoot(a, &kids)
	}
	return Node{Tag: "tfoot", Attrs: a, Kids: kids}
}

func (g Global) applyTfoot(a *TfootAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyTfoot(_ *TfootAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyTfoot(_ *TfootAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TfootAttrs) writeAttrs(sb *strings.Builder)           { writeGlobal(sb, &a.Global) }
