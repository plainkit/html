package html

import "strings"

type EmbedAttrs struct {
	Global GlobalAttrs
	Height string
	Src    string
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

func (o HeightOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o SrcOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o WidthOpt) applyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *EmbedAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
