package html

import "strings"

type TemplateAttrs struct {
	Global                   GlobalAttrs
	Shadowrootclonable       string
	Shadowrootdelegatesfocus string
	Shadowrootmode           string
	Shadowrootserializable   string
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

func (g Global) applyTemplate(a *TemplateAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o TxtOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyTemplate(_ *TemplateAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (o ShadowrootclonableOpt) applyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootclonable = o.v
}
func (o ShadowrootdelegatesfocusOpt) applyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootdelegatesfocus = o.v
}
func (o ShadowrootmodeOpt) applyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootmode = o.v
}
func (o ShadowrootserializableOpt) applyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootserializable = o.v
}

func (a *TemplateAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Shadowrootclonable != "" {
		Attr(sb, "shadowrootclonable", a.Shadowrootclonable)
	}
	if a.Shadowrootdelegatesfocus != "" {
		Attr(sb, "shadowrootdelegatesfocus", a.Shadowrootdelegatesfocus)
	}
	if a.Shadowrootmode != "" {
		Attr(sb, "shadowrootmode", a.Shadowrootmode)
	}
	if a.Shadowrootserializable != "" {
		Attr(sb, "shadowrootserializable", a.Shadowrootserializable)
	}
}
