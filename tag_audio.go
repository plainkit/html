package blox

import "strings"

// Audio
type AudioAttrs struct {
	Global      GlobalAttrs
	Src         string
	Preload     string
	Autoplay    bool
	Loop        bool
	Muted       bool
	Controls    bool
	Crossorigin string
}

type AudioArg interface {
	applyAudio(*AudioAttrs, *[]Component)
}

func defaultAudioAttrs() *AudioAttrs {
	return &AudioAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Audio(args ...AudioArg) Node {
	a := defaultAudioAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyAudio(a, &kids)
	}
	return Node{Tag: "audio", Attrs: a, Kids: kids}
}

func (g Global) applyAudio(a *AudioAttrs, _ *[]Component)         { g.do(&a.Global) }
func (o TxtOpt) applyAudio(_ *AudioAttrs, kids *[]Component)      { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyAudio(_ *AudioAttrs, kids *[]Component)    { *kids = append(*kids, o.c) }
func (o SrcOpt) applyAudio(a *AudioAttrs, _ *[]Component)         { a.Src = o.v }
func (o PreloadOpt) applyAudio(a *AudioAttrs, _ *[]Component)     { a.Preload = o.v }
func (o AutoplayOpt) applyAudio(a *AudioAttrs, _ *[]Component)    { a.Autoplay = true }
func (o LoopOpt) applyAudio(a *AudioAttrs, _ *[]Component)        { a.Loop = true }
func (o MutedOpt) applyAudio(a *AudioAttrs, _ *[]Component)       { a.Muted = true }
func (o ControlsOpt) applyAudio(a *AudioAttrs, _ *[]Component)    { a.Controls = true }
func (o CrossoriginOpt) applyAudio(a *AudioAttrs, _ *[]Component) { a.Crossorigin = o.v }

func (a *AudioAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
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
	if a.Crossorigin != "" {
		attr(sb, "crossorigin", a.Crossorigin)
	}
}
