package html

import "strings"

type RtAttrs struct {
	Global GlobalAttrs
}

type RtArg interface {
	ApplyRt(*RtAttrs, *[]Component)
}

func defaultRtAttrs() *RtAttrs {
	return &RtAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rt(args ...RtArg) Node {
	a := defaultRtAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyRt(a, &kids)
	}
	return Node{Tag: "rt", Attrs: a, Kids: kids}
}

func (g Global) ApplyRt(a *RtAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *RtAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
