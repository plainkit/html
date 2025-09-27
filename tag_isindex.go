package html

import "strings"

type IsindexAttrs struct {
	Global GlobalAttrs
	Prompt string
}

type IsindexArg interface {
	applyIsindex(*IsindexAttrs, *[]Component)
}

func defaultIsindexAttrs() *IsindexAttrs {
	return &IsindexAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Isindex(args ...IsindexArg) Node {
	a := defaultIsindexAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyIsindex(a, &kids)
	}
	return Node{Tag: "isindex", Attrs: a, Kids: kids}
}

func (g Global) applyIsindex(a *IsindexAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o PromptOpt) applyIsindex(a *IsindexAttrs, _ *[]Component) {
	a.Prompt = o.v
}

func (a *IsindexAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Prompt != "" {
		Attr(sb, "prompt", a.Prompt)
	}
}
