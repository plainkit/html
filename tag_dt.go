package html

import "strings"

type DtAttrs struct {
	Global GlobalAttrs
}

type DtArg interface {
	ApplyDt(*DtAttrs, *[]Component)
}

func defaultDtAttrs() *DtAttrs {
	return &DtAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dt(args ...DtArg) Node {
	a := defaultDtAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyDt(a, &kids)
	}

	return Node{Tag: "dt", Attrs: a, Kids: kids}
}

func (g Global) ApplyDt(a *DtAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *DtAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
