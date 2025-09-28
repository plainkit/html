package html

import "strings"

type AudioAttrs struct {
	Global      GlobalAttrs
	Autoplay    bool
	Controls    bool
	Crossorigin string
	Loop        bool
	Muted       bool
	Preload     string
	Src         string
}

type AudioArg interface {
	ApplyAudio(*AudioAttrs, *[]Component)
}

func defaultAudioAttrs() *AudioAttrs {
	return &AudioAttrs{
		Global: GlobalAttrs{
			Style:  "",
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
		ar.ApplyAudio(a, &kids)
	}
	return Node{Tag: "audio", Attrs: a, Kids: kids}
}

func (g Global) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AutoplayOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Autoplay = true
}
func (o ControlsOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Controls = true
}
func (o CrossoriginOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o LoopOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Loop = true
}
func (o MutedOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Muted = true
}
func (o PreloadOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Preload = o.v
}
func (o SrcOpt) ApplyAudio(a *AudioAttrs, _ *[]Component) {
	a.Src = o.v
}

func (a *AudioAttrs) WriteAttrs(sb *strings.Builder) {
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
	if a.Loop {
		BoolAttr(sb, "loop")
	}
	if a.Muted {
		BoolAttr(sb, "muted")
	}
	if a.Preload != "" {
		Attr(sb, "preload", a.Preload)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
}
