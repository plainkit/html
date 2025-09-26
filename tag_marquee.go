package html

import "strings"

type MarqueeAttrs struct {
	Global       GlobalAttrs
	Behavior     string
	Bgcolor      string
	Direction    string
	Height       string
	Hspace       string
	Loop         bool
	Scrollamount string
	Scrolldelay  string
	Truespeed    string
	Vspace       string
	Width        string
}

type MarqueeArg interface {
	applyMarquee(*MarqueeAttrs, *[]Component)
}

func defaultMarqueeAttrs() *MarqueeAttrs {
	return &MarqueeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Marquee(args ...MarqueeArg) Node {
	a := defaultMarqueeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyMarquee(a, &kids)
	}
	return Node{Tag: "marquee", Attrs: a, Kids: kids}
}

func (g Global) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o BehaviorOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Behavior = o.v
}
func (o BgcolorOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Bgcolor = o.v
}
func (o DirectionOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Direction = o.v
}
func (o HeightOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o HspaceOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Hspace = o.v
}
func (o LoopOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Loop = true
}
func (o ScrollamountOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Scrollamount = o.v
}
func (o ScrolldelayOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Scrolldelay = o.v
}
func (o TruespeedOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Truespeed = o.v
}
func (o VspaceOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Vspace = o.v
}
func (o WidthOpt) applyMarquee(a *MarqueeAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *MarqueeAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Behavior != "" {
		Attr(sb, "behavior", a.Behavior)
	}
	if a.Bgcolor != "" {
		Attr(sb, "bgcolor", a.Bgcolor)
	}
	if a.Direction != "" {
		Attr(sb, "direction", a.Direction)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Hspace != "" {
		Attr(sb, "hspace", a.Hspace)
	}
	if a.Loop {
		BoolAttr(sb, "loop")
	}
	if a.Scrollamount != "" {
		Attr(sb, "scrollamount", a.Scrollamount)
	}
	if a.Scrolldelay != "" {
		Attr(sb, "scrolldelay", a.Scrolldelay)
	}
	if a.Truespeed != "" {
		Attr(sb, "truespeed", a.Truespeed)
	}
	if a.Vspace != "" {
		Attr(sb, "vspace", a.Vspace)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
