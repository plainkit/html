package html

import "strings"

// Iframe
type IframeAttrs struct {
	Global          GlobalAttrs
	Src             string
	Srcdoc          string
	Name            string
	Sandbox         string
	Allow           string
	Allowfullscreen bool
	Width           int
	Height          int
	Loading         string
	Referrerpolicy  string
}

type IframeArg interface {
	applyIframe(*IframeAttrs, *[]Component)
}

func defaultIframeAttrs() *IframeAttrs {
	return &IframeAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Iframe(args ...IframeArg) Node {
	a := defaultIframeAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyIframe(a, &kids)
	}
	return Node{Tag: "iframe", Attrs: a, Kids: kids}
}

type SrcdocOpt struct{ v string }
type SandboxOpt struct{ v string }
type AllowOpt struct{ v string }
type AllowfullscreenOpt struct{}
type ReferrerpolicyOpt struct{ v string }

func Srcdoc(v string) SrcdocOpt                 { return SrcdocOpt{v} }
func Sandbox(v string) SandboxOpt               { return SandboxOpt{v} }
func Allow(v string) AllowOpt                   { return AllowOpt{v} }
func Allowfullscreen() AllowfullscreenOpt       { return AllowfullscreenOpt{} }
func Referrerpolicy(v string) ReferrerpolicyOpt { return ReferrerpolicyOpt{v} }

func (g Global) applyIframe(a *IframeAttrs, _ *[]Component)             { g.do(&a.Global) }
func (o TxtOpt) applyIframe(_ *IframeAttrs, kids *[]Component)          { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyIframe(_ *IframeAttrs, kids *[]Component)        { *kids = append(*kids, o.c) }
func (o SrcOpt) applyIframe(a *IframeAttrs, _ *[]Component)             { a.Src = o.v }
func (o SrcdocOpt) applyIframe(a *IframeAttrs, _ *[]Component)          { a.Srcdoc = o.v }
func (o SandboxOpt) applyIframe(a *IframeAttrs, _ *[]Component)         { a.Sandbox = o.v }
func (o AllowOpt) applyIframe(a *IframeAttrs, _ *[]Component)           { a.Allow = o.v }
func (o AllowfullscreenOpt) applyIframe(a *IframeAttrs, _ *[]Component) { a.Allowfullscreen = true }
func (o WidthOpt) applyIframe(a *IframeAttrs, _ *[]Component)           { a.Width = o.v }
func (o HeightOpt) applyIframe(a *IframeAttrs, _ *[]Component)          { a.Height = o.v }
func (o LoadingOpt) applyIframe(a *IframeAttrs, _ *[]Component)         { a.Loading = o.v }
func (o ReferrerpolicyOpt) applyIframe(a *IframeAttrs, _ *[]Component)  { a.Referrerpolicy = o.v }

func (a *IframeAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Srcdoc != "" {
		attr(sb, "srcdoc", a.Srcdoc)
	}
	if a.Name != "" {
		attr(sb, "name", a.Name)
	}
	if a.Sandbox != "" {
		attr(sb, "sandbox", a.Sandbox)
	}
	if a.Allow != "" {
		attr(sb, "allow", a.Allow)
	}
	if a.Allowfullscreen {
		boolAttr(sb, "allowfullscreen")
	}
	if a.Width > 0 {
		attr(sb, "width", itoa(a.Width))
	}
	if a.Height > 0 {
		attr(sb, "height", itoa(a.Height))
	}
	if a.Loading != "" {
		attr(sb, "loading", a.Loading)
	}
	if a.Referrerpolicy != "" {
		attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
}
