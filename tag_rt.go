package html

import "strings"

type RtAttrs struct {
	Global GlobalAttrs
}

type RtArg interface {
	applyRt(*RtAttrs, *[]Component)
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
		ar.applyRt(a, &kids)
	}
	return Node{Tag: "rt", Attrs: a, Kids: kids}
}

func (g Global) applyRt(a *RtAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *RtAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
