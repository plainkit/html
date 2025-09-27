package html

import "strings"

type ObjectAttrs struct {
	Global        GlobalAttrs
	Data          string
	Form          string
	Height        string
	Typemustmatch bool
	Usemap        string
	Width         string
}

type ObjectArg interface {
	applyObject(*ObjectAttrs, *[]Component)
}

func defaultObjectAttrs() *ObjectAttrs {
	return &ObjectAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Object(args ...ObjectArg) Node {
	a := defaultObjectAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyObject(a, &kids)
	}
	return Node{Tag: "object", Attrs: a, Kids: kids}
}

func (g Global) applyObject(a *ObjectAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o FormOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o HeightOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o TypemustmatchOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Typemustmatch = true
}
func (o UsemapOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Usemap = o.v
}
func (o WidthOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ObjectAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Typemustmatch {
		BoolAttr(sb, "typemustmatch")
	}
	if a.Usemap != "" {
		Attr(sb, "usemap", a.Usemap)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
