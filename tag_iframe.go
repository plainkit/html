package html

import "strings"

type IframeAttrs struct {
	Global          GlobalAttrs
	Allow           string
	Allowfullscreen bool
	Height          string
	Loading         string
	Name            string
	Referrerpolicy  string
	Sandbox         string
	Src             string
	Srcdoc          string
	Width           string
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
	return Node{Tag: "iframe", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyIframe(a *IframeAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AllowOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allow = o.v
}
func (o AllowfullscreenOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allowfullscreen = true
}
func (o HeightOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o LoadingOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Loading = o.v
}
func (o NameOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o ReferrerpolicyOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o SandboxOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Sandbox = o.v
}
func (o SrcOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o SrcdocOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Srcdoc = o.v
}
func (o WidthOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Width = o.v
}

func (a *IframeAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Allow != "" {
		Attr(sb, "allow", a.Allow)
	}
	if a.Allowfullscreen {
		BoolAttr(sb, "allowfullscreen")
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Loading != "" {
		Attr(sb, "loading", a.Loading)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Referrerpolicy != "" {
		Attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
	if a.Sandbox != "" {
		Attr(sb, "sandbox", a.Sandbox)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Srcdoc != "" {
		Attr(sb, "srcdoc", a.Srcdoc)
	}
	if a.Width != "" {
		Attr(sb, "width", a.Width)
	}
}
