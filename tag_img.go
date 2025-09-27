package html

import "strings"

type ImgAttrs struct {
	Global         GlobalAttrs
	Alt            string
	Crossorigin    string
	Height         string
	Ismap          bool
	Loading        string
	Longdesc       string
	Referrerpolicy string
	Sizes          string
	Src            string
	Srcset         string
	Usemap         string
	Width          string
}

type ImgArg interface {
	applyImg(*ImgAttrs, *[]Component)
}

func defaultImgAttrs() *ImgAttrs {
	return &ImgAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Img(args ...ImgArg) Node {
	a := defaultImgAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyImg(a, &kids)
	}
	return Node{Tag: "img", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyImg(a *ImgAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AltOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o CrossoriginOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o HeightOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o IsmapOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Ismap = true
}
func (o LoadingOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Loading = o.v
}
func (o LongdescOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Longdesc = o.v
}
func (o ReferrerpolicyOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o SizesOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Sizes = o.v
}
func (o SrcOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o SrcsetOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Srcset = o.v
}
func (o UsemapOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Usemap = o.v
}
func (o WidthOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ImgAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Ismap {
		BoolAttr(sb, "ismap")
	}
	if a.Loading != "" {
		Attr(sb, "loading", a.Loading)
	}
	if a.Longdesc != "" {
		Attr(sb, "longdesc", a.Longdesc)
	}
	if a.Referrerpolicy != "" {
		Attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
	if a.Sizes != "" {
		Attr(sb, "sizes", a.Sizes)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Srcset != "" {
		Attr(sb, "srcset", a.Srcset)
	}
	if a.Usemap != "" {
		Attr(sb, "usemap", a.Usemap)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
