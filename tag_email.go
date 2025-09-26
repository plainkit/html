package html

import "strings"

type EmailAttrs struct {
	Global GlobalAttrs
}

type EmailArg interface {
	applyEmail(*EmailAttrs, *[]Component)
}

func defaultEmailAttrs() *EmailAttrs {
	return &EmailAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Email(args ...EmailArg) Node {
	a := defaultEmailAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyEmail(a, &kids)
	}
	return Node{Tag: "email", Attrs: a, Kids: kids}
}

func (g Global) applyEmail(a *EmailAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *EmailAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
