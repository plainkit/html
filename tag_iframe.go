package html

import "strings"

type IframeAttrs struct {
	Global              GlobalAttrs
	Align               string
	Allow               string
	Allowfullscreen     bool
	Allowpaymentrequest bool
	Allowusermedia      string
	Frameborder         string
	Height              string
	Loading             string
	Longdesc            string
	Marginheight        string
	Marginwidth         string
	Name                string
	Referrerpolicy      string
	Sandbox             string
	Scrolling           string
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
	g.Do(&a.Global)
}

func (o AlignOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Align = o.v
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
func (o AllowusermediaOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Allowusermedia = o.v
}
func (o FrameborderOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Frameborder = o.v
}
func (o HeightOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Height = o.v
}
func (o LoadingOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Loading = o.v
}
func (o LongdescOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Longdesc = o.v
}
func (o MarginheightOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Marginheight = o.v
}
func (o MarginwidthOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Marginwidth = o.v
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
func (o ScrollingOpt) applyIframe(a *IframeAttrs, _ *[]Component) {
	a.Scrolling = o.v
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
	if a.Align != "" {
		Attr(sb, "align", a.Align)
	}
	if a.Allow != "" {
		Attr(sb, "allow", a.Allow)
	}
	if a.Allowfullscreen {
		BoolAttr(sb, "allowfullscreen")
	}
	if a.Allowpaymentrequest {
		BoolAttr(sb, "allowpaymentrequest")
	}
	if a.Allowusermedia != "" {
		Attr(sb, "allowusermedia", a.Allowusermedia)
	}
	if a.Frameborder != "" {
		Attr(sb, "frameborder", a.Frameborder)
	}
	if a.Height != "" {
		Attr(sb, "height", a.Height)
	}
	if a.Loading != "" {
		Attr(sb, "loading", a.Loading)
	}
	if a.Longdesc != "" {
		Attr(sb, "longdesc", a.Longdesc)
	}
	if a.Marginheight != "" {
		Attr(sb, "marginheight", a.Marginheight)
	}
	if a.Marginwidth != "" {
		Attr(sb, "marginwidth", a.Marginwidth)
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
	if a.Scrolling != "" {
		Attr(sb, "scrolling", a.Scrolling)
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
