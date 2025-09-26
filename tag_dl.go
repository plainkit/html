package html

import "strings"

type DlAttrs struct {
	Global  GlobalAttrs
	Compact string
}

type DlArg interface {
	applyDl(*DlAttrs, *[]Component)
}

func defaultDlAttrs() *DlAttrs {
	return &DlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Dl(args ...DlArg) Node {
	a := defaultDlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyDl(a, &kids)
	}
	return Node{Tag: "dl", Attrs: a, Kids: kids}
}

func (g Global) applyDl(a *DlAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o CompactOpt) applyDl(a *DlAttrs, _ *[]Component) {
	a.Compact = o.v
}

func (a *DlAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Compact != "" {
		Attr(sb, "compact", a.Compact)
	}
}
