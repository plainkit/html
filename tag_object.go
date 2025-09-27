package html

import "strings"

type ObjectAttrs struct {
	Global        GlobalAttrs
	Align         string
	Archive       string
	Border        string
	Classid       string
	Codebase      string
	Codetype      string
	Data          string
	Declare       bool
	Form          string
	Height        string
	Hspace        string
	Name          string
	Standby       string
	Type          string
	Typemustmatch bool
	Usemap        string
	Vspace        string
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
	g.Do(&a.Global)
}

func (o AlignOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o ArchiveOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Archive = o.v
}
func (o BorderOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Border = o.v
}
func (o ClassidOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Classid = o.v
}
func (o CodebaseOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Codebase = o.v
}
func (o CodetypeOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Codetype = o.v
}
func (o DeclareOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Declare = true
}
func (o FormOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Form = o.v
}
func (o HeightOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o HspaceOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Hspace = o.v
}
func (o NameOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o StandbyOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Standby = o.v
}
func (o TypeOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o TypemustmatchOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Typemustmatch = true
}
func (o UsemapOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Usemap = o.v
}
func (o VspaceOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Vspace = o.v
}
func (o WidthOpt) applyObject(a *ObjectAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ObjectAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Archive != "" {
		Attr(sb, "archive", a.Archive)
	}
	if a.Border != "" {
		Attr(sb, "border", a.Border)
	}
	if a.Classid != "" {
		Attr(sb, "classid", a.Classid)
	}
	if a.Codebase != "" {
		Attr(sb, "codebase", a.Codebase)
	}
	if a.Codetype != "" {
		Attr(sb, "codetype", a.Codetype)
	}
	if a.Declare {
		BoolAttr(sb, "declare")
	}
	if a.Form != "" {
		Attr(sb, "form", a.Form)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Hspace != "" {
		Attr(sb, "hspace", a.Hspace)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Standby != "" {
		Attr(sb, "standby", a.Standby)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Typemustmatch {
		BoolAttr(sb, "typemustmatch")
	}
	if a.Usemap != "" {
		Attr(sb, "usemap", a.Usemap)
	}
	if a.Vspace != "" {
		Attr(sb, "vspace", a.Vspace)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
