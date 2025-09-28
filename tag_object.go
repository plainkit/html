package html

import "strings"

type ObjectAttrs struct {
	Global GlobalAttrs
	Data   string
	Form   string
	Height string
	Name   string
	Type   string
	Width  string
}

type ObjectArg interface {
	ApplyObject(*ObjectAttrs, *[]Component)
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
		ar.ApplyObject(a, &kids)
	}
	return Node{Tag: "object", Attrs: a, Kids: kids}
}

func (g Global) ApplyObject(a *ObjectAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o FormOpt) ApplyObject(a *ObjectAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o HeightOpt) ApplyObject(a *ObjectAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o NameOpt) ApplyObject(a *ObjectAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o TypeOpt) ApplyObject(a *ObjectAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o WidthOpt) ApplyObject(a *ObjectAttrs, _ *[]Component) {
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
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
