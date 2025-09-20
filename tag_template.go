package html

import "strings"

// Template
type TemplateAttrs struct {
	Global GlobalAttrs
}

type TemplateArg interface {
	applyTemplate(*TemplateAttrs, *[]Component)
}

func defaultTemplateAttrs() *TemplateAttrs {
	return &TemplateAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Template(args ...TemplateArg) Node {
	a := defaultTemplateAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTemplate(a, &kids)
	}
	return Node{Tag: "template", Attrs: a, Kids: kids}
}

func (g Global) applyTemplate(a *TemplateAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *TemplateAttrs) writeAttrs(sb *strings.Builder)              { writeGlobal(sb, &a.Global) }
