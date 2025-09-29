package html

import "strings"

type BlockquoteAttrs struct {
	Global GlobalAttrs
	Cite   string
}

type BlockquoteArg interface {
	ApplyBlockquote(*BlockquoteAttrs, *[]Component)
}

func defaultBlockquoteAttrs() *BlockquoteAttrs {
	return &BlockquoteAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Blockquote(args ...BlockquoteArg) Node {
	a := defaultBlockquoteAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyBlockquote(a, &kids)
	}

	return Node{Tag: "blockquote", Attrs: a, Kids: kids}
}

func (g Global) ApplyBlockquote(a *BlockquoteAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o CiteOpt) ApplyBlockquote(a *BlockquoteAttrs, _ *[]Component) {
	a.Cite = o.v
}

func (a *BlockquoteAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.Cite != "" {
		Attr(sb, "cite", a.Cite)
	}
}
