package html

import "strings"

type DirAttrs struct {
	Global  GlobalAttrs
	Compact string
}

type DirArg interface {
	applyDir(*DirAttrs, *[]Component)
}

func defaultDirAttrs() *DirAttrs {
	return &DirAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dir(args ...DirArg) Node {
	a := defaultDirAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDir(a, &kids)
	}
	return Node{Tag: "dir", Attrs: a, Kids: kids}
}

func (g Global) applyDir(a *DirAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o CompactOpt) applyDir(a *DirAttrs, _ *[]Component) {
	a.Compact = o.v
}

func (a *DirAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Compact != "" {
		Attr(sb, "compact", a.Compact)
	}
}
