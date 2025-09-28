package html

import "strings"

type TemplateAttrs struct {
	Global                          GlobalAttrs
	Shadowrootclonable              bool
	Shadowrootcustomelementregistry bool
	Shadowrootdelegatesfocus        bool
	Shadowrootmode                  string
	Shadowrootserializable          bool
}

type TemplateArg interface {
	ApplyTemplate(*TemplateAttrs, *[]Component)
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
		ar.ApplyTemplate(a, &kids)
	}
	return Node{Tag: "template", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ShadowrootclonableOpt) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootclonable = true
}
func (o ShadowrootcustomelementregistryOpt) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootcustomelementregistry = true
}
func (o ShadowrootdelegatesfocusOpt) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootdelegatesfocus = true
}
func (o ShadowrootmodeOpt) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootmode = o.v
}
func (o ShadowrootserializableOpt) ApplyTemplate(a *TemplateAttrs, _ *[]Component) {
	a.Shadowrootserializable = true
}

func (a *TemplateAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Shadowrootclonable {
		BoolAttr(sb, "shadowrootclonable")
	}
	if a.Shadowrootcustomelementregistry {
		BoolAttr(sb, "shadowrootcustomelementregistry")
	}
	if a.Shadowrootdelegatesfocus {
		BoolAttr(sb, "shadowrootdelegatesfocus")
	}
	if a.Shadowrootmode != "" {
		Attr(sb, "shadowrootmode", a.Shadowrootmode)
	}
	if a.Shadowrootserializable {
		BoolAttr(sb, "shadowrootserializable")
	}
}
