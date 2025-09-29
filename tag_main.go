package html

import "strings"

type MainAttrs struct {
	Global GlobalAttrs
}

type MainArg interface {
	ApplyMain(*MainAttrs, *[]Component)
}

func defaultMainAttrs() *MainAttrs {
	return &MainAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Main(args ...MainArg) Node {
	a := defaultMainAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyMain(a, &kids)
	}

	return Node{Tag: "main", Attrs: a, Kids: kids}
}

func (g Global) ApplyMain(a *MainAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *MainAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
