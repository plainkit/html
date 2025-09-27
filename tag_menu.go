package html

import "strings"

type MenuAttrs struct {
	Global GlobalAttrs
}

type MenuArg interface {
	applyMenu(*MenuAttrs, *[]Component)
}

func defaultMenuAttrs() *MenuAttrs {
	return &MenuAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Menu(args ...MenuArg) Node {
	a := defaultMenuAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMenu(a, &kids)
	}
	return Node{Tag: "menu", Attrs: a, Kids: kids}
}

func (g Global) applyMenu(a *MenuAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *MenuAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
