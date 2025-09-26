package html

import "strings"

type ImgAttrs struct {
	Global                            GlobalAttrs
	Align                             string
	Alt                               string
	AspectRatioComputedFromAttributes string
	Attributionsrc                    string
	Border                            string
	Crossorigin                       string
	Decoding                          string
	Fetchpriority                     string
	Height                            string
	Hspace                            string
	Ismap                             string
	Loading                           string
	Longdesc                          string
	Name                              string
	Referrerpolicy                    string
	Sizes                             string
	Src                               string
	Srcset                            string
	Usemap                            string
	Vspace                            string
	Width                             string
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

func (o AlignOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Align = o.v
}
func (o AltOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o AspectRatioComputedFromAttributesOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.AspectRatioComputedFromAttributes = o.v
}
func (o AttributionsrcOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Attributionsrc = o.v
}
func (o BorderOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Border = o.v
}
func (o CrossoriginOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DecodingOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Decoding = o.v
}
func (o FetchpriorityOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o HeightOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o HspaceOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Hspace = o.v
}
func (o IsmapOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Ismap = o.v
}
func (o LoadingOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Loading = o.v
}
func (o LongdescOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Longdesc = o.v
}
func (o NameOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Name = o.v
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
func (o VspaceOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Vspace = o.v
}
func (o WidthOpt) applyImg(a *ImgAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *ImgAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
	}
	if a.AspectRatioComputedFromAttributes != "" {
		Attr(sb, "aspect_ratio_computed_from_attributes", a.AspectRatioComputedFromAttributes)
	}
	if a.Attributionsrc != "" {
		Attr(sb, "attributionsrc", a.Attributionsrc)
	}
	if a.Border != "" {
		Attr(sb, "border", a.Border)
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
	if a.Hspace != "" {
		Attr(sb, "hspace", a.Hspace)
	}
	if a.Ismap != "" {
		Attr(sb, "ismap", a.Ismap)
	}
	if a.Loading != "" {
		Attr(sb, "loading", a.Loading)
	}
	if a.Longdesc != "" {
		Attr(sb, "longdesc", a.Longdesc)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
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
	if a.Vspace != "" {
		Attr(sb, "vspace", a.Vspace)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
