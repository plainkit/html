package html

import "strings"

// Video
type VideoAttrs struct {
	Global      GlobalAttrs
	Src         string
	Poster      string
	Preload     string
	Autoplay    bool
	Loop        bool
	Muted       bool
	Controls    bool
	Width       int
	Height      int
	Crossorigin string
}

type VideoArg interface {
	applyVideo(*VideoAttrs, *[]Component)
}

func defaultVideoAttrs() *VideoAttrs {
	return &VideoAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
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

// Video-specific options
type PosterOpt struct{ v string }
type PreloadOpt struct{ v string }
type AutoplayOpt struct{}
type LoopOpt struct{}
type MutedOpt struct{}
type ControlsOpt struct{}

func Poster(v string) PosterOpt   { return PosterOpt{v} }
func Preload(v string) PreloadOpt { return PreloadOpt{v} }
func Autoplay() AutoplayOpt       { return AutoplayOpt{} }
func Loop() LoopOpt               { return LoopOpt{} }
func Muted() MutedOpt             { return MutedOpt{} }
func Controls() ControlsOpt       { return ControlsOpt{} }

func (g Global) applyVideo(a *VideoAttrs, _ *[]Component)         { g.do(&a.Global) }
func (o TxtOpt) applyVideo(_ *VideoAttrs, kids *[]Component)      { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyVideo(_ *VideoAttrs, kids *[]Component)    { *kids = append(*kids, o.c) }
func (o SrcOpt) applyVideo(a *VideoAttrs, _ *[]Component)         { a.Src = o.v }
func (o PosterOpt) applyVideo(a *VideoAttrs, _ *[]Component)      { a.Poster = o.v }
func (o PreloadOpt) applyVideo(a *VideoAttrs, _ *[]Component)     { a.Preload = o.v }
func (o AutoplayOpt) applyVideo(a *VideoAttrs, _ *[]Component)    { a.Autoplay = true }
func (o LoopOpt) applyVideo(a *VideoAttrs, _ *[]Component)        { a.Loop = true }
func (o MutedOpt) applyVideo(a *VideoAttrs, _ *[]Component)       { a.Muted = true }
func (o ControlsOpt) applyVideo(a *VideoAttrs, _ *[]Component)    { a.Controls = true }
func (o WidthOpt) applyVideo(a *VideoAttrs, _ *[]Component)       { a.Width = o.v }
func (o HeightOpt) applyVideo(a *VideoAttrs, _ *[]Component)      { a.Height = o.v }
func (o CrossoriginOpt) applyVideo(a *VideoAttrs, _ *[]Component) { a.Crossorigin = o.v }

func (a *VideoAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Poster != "" {
		attr(sb, "poster", a.Poster)
	}
	if a.Preload != "" {
		attr(sb, "preload", a.Preload)
	}
	if a.Autoplay {
		boolAttr(sb, "autoplay")
	}
	if a.Loop {
		boolAttr(sb, "loop")
	}
	if a.Muted {
		boolAttr(sb, "muted")
	}
	if a.Controls {
		boolAttr(sb, "controls")
	}
	if a.Width > 0 {
		attr(sb, "width", itoa(a.Width))
	}
	if a.Height > 0 {
		attr(sb, "height", itoa(a.Height))
	}
	if a.Crossorigin != "" {
		attr(sb, "crossorigin", a.Crossorigin)
	}
}
