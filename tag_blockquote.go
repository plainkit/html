package html

import "strings"

type BlockquoteAttrs struct {
	Global GlobalAttrs
	Cite   string
}

type BlockquoteArg interface {
	applyBlockquote(*BlockquoteAttrs, *[]Component)
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
		ar.applyBlockquote(a, &kids)
	}
	return Node{Tag: "blockquote", Attrs: a, Kids: kids}
}

// Tag-specific options
type CiteOpt struct{ v string }

func Cite(v string) CiteOpt { return CiteOpt{v} }

// Global option glue
func (g Global) applyBlockquote(a *BlockquoteAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyBlockquote(_ *BlockquoteAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Tag-specific option glue
func (o CiteOpt) applyBlockquote(a *BlockquoteAttrs, _ *[]Component) {
	a.Cite = o.v
}

// Attrs writer implementation
func (a *BlockquoteAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Cite != "" {
		attr(sb, "cite", a.Cite)
	}
}
