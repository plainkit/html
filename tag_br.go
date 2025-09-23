package html

import "strings"

type BrAttrs struct {
	Global GlobalAttrs
	Clear  string
}

type BrArg interface {
	applyBr(*BrAttrs, *[]Component)
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
		ar.applyBr(a, &kids)
	}
	return Node{Tag: "br", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyBr(a *BrAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyBr(_ *BrAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyBr(_ *BrAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o ClearOpt) applyBr(a *BrAttrs, _ *[]Component) {
	a.Clear = o.v
}

func (a *BrAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Clear != "" {
		Attr(sb, "clear", a.Clear)
	}
}
