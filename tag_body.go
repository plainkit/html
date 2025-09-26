package html

import "strings"

type BodyAttrs struct {
	Global     GlobalAttrs
	Alink      string
	Background string
	Bgcolor    string
	Link       string
	Text       string
	Vlink      string
}

type BodyArg interface {
	applyBody(*BodyAttrs, *[]Component)
}

func defaultBodyAttrs() *BodyAttrs {
	return &BodyAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Body(args ...BodyArg) Node {
	a := defaultBodyAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyBody(a, &kids)
	}
	return Node{Tag: "body", Attrs: a, Kids: kids}
}

func (g Global) applyBody(a *BodyAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlinkOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Alink = o.v
}
func (o BackgroundOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Background = o.v
}
func (o BgcolorOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o LinkOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Link = o.v
}
func (o TextOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Text = o.v
}
func (o VlinkOpt) applyBody(a *BodyAttrs, _ *[]Component) {
	a.Vlink = o.v
}

func (a *BodyAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Alink != "" {
		Attr(sb, "alink", a.Alink)
	}
	if a.Background != "" {
		Attr(sb, "background", a.Background)
	}
	if a.Bgcolor != "" {
		Attr(sb, "bgcolor", a.Bgcolor)
	}
	if a.Link != "" {
		Attr(sb, "link", a.Link)
	}
	if a.Text != "" {
		Attr(sb, "text", a.Text)
	}
	if a.Vlink != "" {
		Attr(sb, "vlink", a.Vlink)
	}
}
