package html

import "strings"

type VideoAttrs struct {
	Global      GlobalAttrs
	Autoplay    bool
	Controls    bool
	Crossorigin string
	Height      string
	Loop        bool
	Muted       bool
	Playsinline bool
	Poster      string
	Preload     string
	Src         string
	Width       string
}

type VideoArg interface {
	ApplyVideo(*VideoAttrs, *[]Component)
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
		ar.ApplyVideo(a, &kids)
	}
	return Node{Tag: "video", Attrs: a, Kids: kids}
}

func (g Global) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AutoplayOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Autoplay = true
}
func (o ControlsOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Controls = true
}
func (o CrossoriginOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o HeightOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o LoopOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Loop = true
}
func (o MutedOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Muted = true
}
func (o PlaysinlineOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Playsinline = true
}
func (o PosterOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Poster = o.v
}
func (o PreloadOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Preload = o.v
}
func (o SrcOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o WidthOpt) ApplyVideo(a *VideoAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *VideoAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Autoplay {
		BoolAttr(sb, "autoplay")
	}
	if a.Controls {
		BoolAttr(sb, "controls")
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
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
	if a.Playsinline {
		BoolAttr(sb, "playsinline")
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
