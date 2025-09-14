package blox

import "strings"

type ImgAttrs struct {
	Global   GlobalAttrs
	Src      string
	Alt      string
	Width    int
	Height   int
	Decoding string // "auto"|"async"
	Loading  string // "lazy"|"eager"
}

type ImgArg interface {
	applyImg(*ImgAttrs)
}

func defaultImgAttrs() *ImgAttrs {
	return &ImgAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
		Decoding: "auto",
		Loading:  "lazy",
	}
}

func Img(args ...ImgArg) Node {
	a := defaultImgAttrs()
	for _, ar := range args {
		ar.applyImg(a)
	}
	return Node{Tag: "img", Attrs: a, Void: true}
}

// Tag-specific options
type SrcOpt struct {
	v string
}

type AltOpt struct {
	v string
}

type WidthOpt struct {
	v int
}

type HeightOpt struct {
	v int
}

type DecodingOpt struct {
	v string
}

type LoadingOpt struct {
	v string
}

func Src(v string) SrcOpt {
	return SrcOpt{v}
}

func Alt(v string) AltOpt {
	return AltOpt{v}
}

func Width(v int) WidthOpt {
	return WidthOpt{v}
}

func Height(v int) HeightOpt {
	return HeightOpt{v}
}

func Decoding(v string) DecodingOpt {
	return DecodingOpt{v}
}

func Loading(v string) LoadingOpt {
	return LoadingOpt{v}
}

// Global option glue
func (g Global) applyImg(a *ImgAttrs) {
	g.do(&a.Global)
}

// Tag-specific option glue
func (o SrcOpt) applyImg(a *ImgAttrs) {
	a.Src = o.v
}

func (o AltOpt) applyImg(a *ImgAttrs) {
	a.Alt = o.v
}

func (o WidthOpt) applyImg(a *ImgAttrs) {
	a.Width = o.v
}

func (o HeightOpt) applyImg(a *ImgAttrs) {
	a.Height = o.v
}

func (o DecodingOpt) applyImg(a *ImgAttrs) {
	a.Decoding = o.v
}

func (o LoadingOpt) applyImg(a *ImgAttrs) {
	a.Loading = o.v
}

// Attrs writer implementation
func (a *ImgAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Alt != "" {
		attr(sb, "alt", a.Alt)
	}
	if a.Width > 0 {
		attr(sb, "width", itoa(a.Width))
	}
	if a.Height > 0 {
		attr(sb, "height", itoa(a.Height))
	}
	if a.Decoding != "" {
		attr(sb, "decoding", a.Decoding)
	}
	if a.Loading != "" {
		attr(sb, "loading", a.Loading)
	}
}
