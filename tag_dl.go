package html

import "strings"

type DlAttrs struct {
	Global GlobalAttrs
}

type DlArg interface {
	ApplyDl(*DlAttrs, *[]Component)
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
		ar.ApplyDl(a, &kids)
	}

	return Node{Tag: "dl", Attrs: a, Kids: kids}
}

func (g Global) ApplyDl(a *DlAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *DlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
