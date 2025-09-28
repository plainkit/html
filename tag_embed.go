package html

import "strings"

type EmbedAttrs struct {
	Global GlobalAttrs
	Height string
	Src    string
	Type   string
	Width  string
}

type EmbedArg interface {
	ApplyEmbed(*EmbedAttrs, *[]Component)
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
		ar.ApplyEmbed(a, &kids)
	}
	return Node{Tag: "embed", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyEmbed(a *EmbedAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o HeightOpt) ApplyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o SrcOpt) ApplyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o TypeOpt) ApplyEmbed(a *EmbedAttrs, _ *[]Component) {
	a.Type = o.v
}
func (o WidthOpt) ApplyEmbed(a *EmbedAttrs, _ *[]Component) {
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
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
