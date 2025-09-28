package html

import "strings"

type HgroupAttrs struct {
	Global GlobalAttrs
}

type HgroupArg interface {
	ApplyHgroup(*HgroupAttrs, *[]Component)
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
		ar.ApplyHgroup(a, &kids)
	}
	return Node{Tag: "hgroup", Attrs: a, Kids: kids}
}

func (g Global) ApplyHgroup(a *HgroupAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *HgroupAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
