package html

import "strings"

type PictureAttrs struct {
	Global GlobalAttrs
	Height string
	Width  string
}

type PictureArg interface {
	applyPicture(*PictureAttrs, *[]Component)
}

func defaultPictureAttrs() *PictureAttrs {
	return &PictureAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Picture(args ...PictureArg) Node {
	a := defaultPictureAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyPicture(a, &kids)
	}
	return Node{Tag: "picture", Attrs: a, Kids: kids}
}

func (g Global) applyPicture(a *PictureAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o HeightOpt) applyPicture(a *PictureAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o WidthOpt) applyPicture(a *PictureAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *PictureAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
