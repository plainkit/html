package blox

import "strings"

// Track (void)
type TrackAttrs struct {
	Global  GlobalAttrs
	Kind    string
	Src     string
	Srclang string
	Label   string
	Default bool
}

type TrackArg interface {
	applyTrack(*TrackAttrs)
}

func defaultTrackAttrs() *TrackAttrs {
	return &TrackAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Track(args ...TrackArg) Node {
	a := defaultTrackAttrs()
	for _, ar := range args {
		ar.applyTrack(a)
	}
	return Node{Tag: "track", Attrs: a, Void: true}
}

// Track-specific options
type KindOpt struct{ v string }
type SrclangOpt struct{ v string }
type LabelOpt struct{ v string }
type DefaultOpt struct{}

func Kind(v string) KindOpt       { return KindOpt{v} }
func Srclang(v string) SrclangOpt { return SrclangOpt{v} }
func Label(v string) LabelOpt     { return LabelOpt{v} }
func Default() DefaultOpt         { return DefaultOpt{} }

func (g Global) applyTrack(a *TrackAttrs)     { g.do(&a.Global) }
func (o KindOpt) applyTrack(a *TrackAttrs)    { a.Kind = o.v }
func (o SrcOpt) applyTrack(a *TrackAttrs)     { a.Src = o.v }
func (o SrclangOpt) applyTrack(a *TrackAttrs) { a.Srclang = o.v }
func (o LabelOpt) applyTrack(a *TrackAttrs)   { a.Label = o.v }
func (o DefaultOpt) applyTrack(a *TrackAttrs) { a.Default = true }

func (a *TrackAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Kind != "" {
		attr(sb, "kind", a.Kind)
	}
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Srclang != "" {
		attr(sb, "srclang", a.Srclang)
	}
	if a.Label != "" {
		attr(sb, "label", a.Label)
	}
	if a.Default {
		boolAttr(sb, "default")
	}
}
