package html

import "strings"

type UlAttrs struct {
	Global  GlobalAttrs
	Compact string
	Type    string
}

type UlArg interface {
	applyUl(*UlAttrs, *[]Component)
}

func defaultUlAttrs() *UlAttrs {
	return &UlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ul(args ...UlArg) Node {
	a := defaultUlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyUl(a, &kids)
	}
	return Node{Tag: "ul", Attrs: a, Kids: kids}
}

func (g Global) applyUl(a *UlAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o CompactOpt) applyUl(a *UlAttrs, _ *[]Component) {
	a.Compact = o.v
}
func (o TypeOpt) applyUl(a *UlAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *UlAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Compact != "" {
		Attr(sb, "compact", a.Compact)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
