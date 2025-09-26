package html

import "strings"

type EmbedAttrs struct {
	Global GlobalAttrs
	Align  string
	Height string
	Name   string
	Src    string
	Type   string
	Width  string
}

type EmbedArg interface {
	applyEmbed(*EmbedAttrs, *[]Component)
}

func defaultEmbedAttrs() *EmbedAttrs {
	return &EmbedAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Embed(args ...EmbedArg) Node {
	a := defaultEmbedAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyEmbed(a, &kids)
	}
	return Node{Tag: "embed", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AlignOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o HeightOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o NameOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o SrcOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o TypeOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o WidthOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *EmbedAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
