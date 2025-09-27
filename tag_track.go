package html

import "strings"

type TrackAttrs struct {
	Global  GlobalAttrs
	Default bool
	Kind    string
	Label   string
	Src     string
	Srclang string
}

type TrackArg interface {
	applyTrack(*TrackAttrs, *[]Component)
}

func defaultTrackAttrs() *TrackAttrs {
	return &TrackAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Track(args ...TrackArg) Node {
	a := defaultTrackAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyTrack(a, &kids)
	}
	return Node{Tag: "track", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyTrack(a *TrackAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DefaultOpt) applyTrack(a *TrackAttrs, _ *[]Component) {
	a.Default = true
}
func (o KindOpt) applyTrack(a *TrackAttrs, _ *[]Component) {
	a.Kind = o.v
}
func (o LabelOpt) applyTrack(a *TrackAttrs, _ *[]Component) {
	a.Label = o.v
}
func (o SrcOpt) applyTrack(a *TrackAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o SrclangOpt) applyTrack(a *TrackAttrs, _ *[]Component) {
	a.Srclang = o.v
}

func (a *TrackAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Default {
		BoolAttr(sb, "default")
	}
	if a.Kind != "" {
		Attr(sb, "kind", a.Kind)
	}
	if a.Label != "" {
		Attr(sb, "label", a.Label)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Srclang != "" {
		Attr(sb, "srclang", a.Srclang)
	}
}
