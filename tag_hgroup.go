package html

import "strings"

type HgroupAttrs struct {
	Global GlobalAttrs
}

type HgroupArg interface {
	applyHgroup(*HgroupAttrs, *[]Component)
}

func defaultHgroupAttrs() *HgroupAttrs {
	return &HgroupAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Hgroup(args ...HgroupArg) Node {
	a := defaultHgroupAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHgroup(a, &kids)
	}
	return Node{Tag: "hgroup", Attrs: a, Kids: kids}
}

func (g Global) applyHgroup(a *HgroupAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *HgroupAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
