package html

import "strings"

type AppletAttrs struct {
	Global   GlobalAttrs
	Align    string
	Alt      string
	Archive  string
	Code     string
	Codebase string
	Height   string
	Hspace   string
	Name     string
	Object   string
	Vspace   string
	Width    string
}

type AppletArg interface {
	applyApplet(*AppletAttrs, *[]Component)
}

func defaultAppletAttrs() *AppletAttrs {
	return &AppletAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Applet(args ...AppletArg) Node {
	a := defaultAppletAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyApplet(a, &kids)
	}
	return Node{Tag: "applet", Attrs: a, Kids: kids}
}

func (g Global) applyApplet(a *AppletAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o AltOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o ArchiveOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Archive = o.v
}
func (o CodeOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Code = o.v
}
func (o CodebaseOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Codebase = o.v
}
func (o HeightOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o HspaceOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Hspace = o.v
}
func (o NameOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o ObjectOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Object = o.v
}
func (o VspaceOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Vspace = o.v
}
func (o WidthOpt) applyApplet(a *AppletAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *AppletAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
	}
	if a.Archive != "" {
		Attr(sb, "archive", a.Archive)
	}
	if a.Code != "" {
		Attr(sb, "code", a.Code)
	}
	if a.Codebase != "" {
		Attr(sb, "codebase", a.Codebase)
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
	if a.Object != "" {
		Attr(sb, "object", a.Object)
	}
	if a.Vspace != "" {
		Attr(sb, "vspace", a.Vspace)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
