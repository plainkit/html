package html

import "strings"

type VideoAttrs struct {
	Global                            GlobalAttrs
	AspectRatioComputedFromAttributes string
	Autoplay                          bool
	Controls                          bool
	Controlslist                      string
	Crossorigin                       string
	Disablepictureinpicture           string
	Disableremoteplayback             string
	Height                            string
	Loop                              bool
	Muted                             bool
	Playsinline                       string
	Poster                            string
	Preload                           string
	Src                               string
	Width                             string
}

type VideoArg interface {
	applyVideo(*VideoAttrs, *[]Component)
}

func defaultVideoAttrs() *VideoAttrs {
	return &VideoAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Video(args ...VideoArg) Node {
	a := defaultVideoAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyVideo(a, &kids)
	}
	return Node{Tag: "video", Attrs: a, Kids: kids}
}

func (g Global) applyVideo(a *VideoAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AspectRatioComputedFromAttributesOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.AspectRatioComputedFromAttributes = o.v
}
func (o AutoplayOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Autoplay = true
}
func (o ControlsOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Controls = true
}
func (o ControlslistOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Controlslist = o.v
}
func (o CrossoriginOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DisablepictureinpictureOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Disablepictureinpicture = o.v
}
func (o DisableremoteplaybackOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Disableremoteplayback = o.v
}
func (o HeightOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o LoopOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Loop = true
}
func (o MutedOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Muted = true
}
func (o PlaysinlineOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Playsinline = o.v
}
func (o PosterOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Poster = o.v
}
func (o PreloadOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Preload = o.v
}
func (o SrcOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o WidthOpt) applyVideo(a *VideoAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *VideoAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.AspectRatioComputedFromAttributes != "" {
		Attr(sb, "aspect_ratio_computed_from_attributes", a.AspectRatioComputedFromAttributes)
	}
	if a.Autoplay {
		BoolAttr(sb, "autoplay")
	}
	if a.Controls {
		BoolAttr(sb, "controls")
	}
	if a.Controlslist != "" {
		Attr(sb, "controlslist", a.Controlslist)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Disablepictureinpicture != "" {
		Attr(sb, "disablepictureinpicture", a.Disablepictureinpicture)
	}
	if a.Disableremoteplayback != "" {
		Attr(sb, "disableremoteplayback", a.Disableremoteplayback)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Loop {
		BoolAttr(sb, "loop")
	}
	if a.Muted {
		BoolAttr(sb, "muted")
	}
	if a.Playsinline != "" {
		Attr(sb, "playsinline", a.Playsinline)
	}
	if a.Poster != "" {
		Attr(sb, "poster", a.Poster)
	}
	if a.Preload != "" {
		Attr(sb, "preload", a.Preload)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
