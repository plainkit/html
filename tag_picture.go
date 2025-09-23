package html

import "strings"

type PictureAttrs struct {
	Global GlobalAttrs
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
	g.do(&a.Global)
}

func (o TxtOpt) applyPicture(_ *PictureAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyPicture(_ *PictureAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

func (a *PictureAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
