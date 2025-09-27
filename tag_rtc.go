package html

import "strings"

type RtcAttrs struct {
	Global GlobalAttrs
}

type RtcArg interface {
	applyRtc(*RtcAttrs, *[]Component)
}

func defaultRtcAttrs() *RtcAttrs {
	return &RtcAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Rtc(args ...RtcArg) Node {
	a := defaultRtcAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyRtc(a, &kids)
	}
	return Node{Tag: "rtc", Attrs: a, Kids: kids}
}

func (g Global) applyRtc(a *RtcAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *RtcAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
