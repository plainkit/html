package html

import "strings"

type DfnAttrs struct {
	Global GlobalAttrs
}

type DfnArg interface {
	applyDfn(*DfnAttrs, *[]Component)
}

func defaultDfnAttrs() *DfnAttrs {
	return &DfnAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dfn(args ...DfnArg) Node {
	a := defaultDfnAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDfn(a, &kids)
	}
	return Node{Tag: "dfn", Attrs: a, Kids: kids}
}

func (g Global) applyDfn(a *DfnAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *DfnAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
