package html

import "strings"

type AudioAttrs struct {
	Global   GlobalAttrs
	Autoplay bool
	Controls bool
	Loop     bool
	Muted    bool
	Preload  string
	Src      string
}

type AudioArg interface {
	applyAudio(*AudioAttrs, *[]Component)
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
		ar.applyAudio(a, &kids)
	}
	return Node{Tag: "audio", Attrs: a, Kids: kids}
}

func (g Global) applyAudio(a *AudioAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AutoplayOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
	a.Autoplay = true
}
func (o ControlsOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
	a.Controls = true
}
func (o LoopOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
	a.Loop = true
}
func (o MutedOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
	a.Muted = true
}
func (o PreloadOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
	a.Preload = o.v
}
func (o SrcOpt) applyAudio(a *AudioAttrs, _ *[]Component) {
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
