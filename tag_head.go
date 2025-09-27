package html

import "strings"

type HeadAttrs struct {
	Global  GlobalAttrs
	Profile string
}

type HeadArg interface {
	applyHead(*HeadAttrs, *[]Component)
}

func defaultHeadAttrs() *HeadAttrs {
	return &HeadAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Head(args ...HeadArg) Node {
	a := defaultHeadAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHead(a, &kids)
	}
	return Node{Tag: "head", Attrs: a, Kids: kids}
}

func (g Global) applyHead(a *HeadAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ProfileOpt) applyHead(a *HeadAttrs, _ *[]Component) {
	a.Profile = o.v
}

func (a *HeadAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Profile != "" {
		Attr(sb, "profile", a.Profile)
	}
}
