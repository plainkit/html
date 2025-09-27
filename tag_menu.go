package html

import "strings"

type MenuAttrs struct {
	Global  GlobalAttrs
	Compact bool
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
	g.Do(&a.Global)
}

func (o CompactOpt) applyMenu(a *MenuAttrs, _ *[]Component) {
	a.Compact = true
}

func (a *MenuAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Compact {
		BoolAttr(sb, "compact")
	}
}
