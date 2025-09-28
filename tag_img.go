package html

import "strings"

type ImgAttrs struct {
	Global         GlobalAttrs
	Alt            string
	Crossorigin    string
	Decoding       string
	Fetchpriority  string
	Height         string
	Ismap          bool
	Loading        string
	Referrerpolicy string
	Sizes          string
	Src            string
	Srcset         string
	Usemap         string
	Width          string
}

type ImgArg interface {
	ApplyImg(*ImgAttrs, *[]Component)
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
		ar.ApplyImg(a, &kids)
	}
	return Node{Tag: "img", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyImg(a *ImgAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AltOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o CrossoriginOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DecodingOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Decoding = o.v
}
func (o FetchpriorityOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o HeightOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o IsmapOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Ismap = true
}
func (o LoadingOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Loading = o.v
}
func (o ReferrerpolicyOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o SizesOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Sizes = o.v
}
func (o SrcOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o SrcsetOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Srcset = o.v
}
func (o UsemapOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Usemap = o.v
}
func (o WidthOpt) ApplyImg(a *ImgAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ImgAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Decoding != "" {
		Attr(sb, "decoding", a.Decoding)
	}
	if a.Fetchpriority != "" {
		Attr(sb, "fetchpriority", a.Fetchpriority)
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
