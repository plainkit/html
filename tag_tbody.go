package html

import "strings"

type TbodyAttrs struct {
	Global GlobalAttrs
}

type TbodyArg interface {
	applyTbody(*TbodyAttrs, *[]Component)
}

func defaultTbodyAttrs() *TbodyAttrs {
	return &TbodyAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Tbody(args ...TbodyArg) Node {
	a := defaultTbodyAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTbody(a, &kids)
	}
	return Node{Tag: "tbody", Attrs: a, Kids: kids}
}

func (g Global) applyTbody(a *TbodyAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *TbodyAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
