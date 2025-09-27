package html

import "strings"

type IframeAttrs struct {
	Global              GlobalAttrs
	Allow               string
	Allowfullscreen     bool
	Allowpaymentrequest bool
	Height              string
	Referrerpolicy      string
	Sandbox             string
	Src                 string
	Srcdoc              string
	Width               string
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

func (g Global) applyIframe(a *IframeAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AllowOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allow = o.v
}
func (o AllowfullscreenOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allowfullscreen = true
}
func (o AllowpaymentrequestOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allowpaymentrequest = true
}
func (o HeightOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Height = o.v
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

func (a *IframeAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Allow != "" {
		Attr(sb, "allow", a.Allow)
	}
	if a.Allowfullscreen {
		BoolAttr(sb, "allowfullscreen")
	}
	if a.Allowpaymentrequest {
		BoolAttr(sb, "allowpaymentrequest")
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
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
