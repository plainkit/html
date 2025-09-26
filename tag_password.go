package html

import "strings"

type PasswordAttrs struct {
	Global GlobalAttrs
}

type PasswordArg interface {
	applyPassword(*PasswordAttrs, *[]Component)
}

func defaultPasswordAttrs() *PasswordAttrs {
	return &PasswordAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Password(args ...PasswordArg) Node {
	a := defaultPasswordAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPassword(a, &kids)
	}
	return Node{Tag: "password", Attrs: a, Kids: kids}
}

func (g Global) applyPassword(a *PasswordAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *PasswordAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
