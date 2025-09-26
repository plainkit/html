package html

import "strings"

type UrlAttrs struct {
	Global GlobalAttrs
}

type UrlArg interface {
	applyUrl(*UrlAttrs, *[]Component)
}

func defaultUrlAttrs() *UrlAttrs {
	return &UrlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Url(args ...UrlArg) Node {
	a := defaultUrlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyUrl(a, &kids)
	}
	return Node{Tag: "url", Attrs: a, Kids: kids}
}

func (g Global) applyUrl(a *UrlAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *UrlAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
