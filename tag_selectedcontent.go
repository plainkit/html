package html

import "strings"

type SelectedcontentAttrs struct {
	Global GlobalAttrs
}

type SelectedcontentArg interface {
	ApplySelectedcontent(*SelectedcontentAttrs, *[]Component)
}

func defaultSelectedcontentAttrs() *SelectedcontentAttrs {
	return &SelectedcontentAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Selectedcontent(args ...SelectedcontentArg) Node {
	a := defaultSelectedcontentAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplySelectedcontent(a, &kids)
	}
	return Node{Tag: "selectedcontent", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplySelectedcontent(a *SelectedcontentAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *SelectedcontentAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
