package html

import "strings"

type AAttrs struct {
	Global         GlobalAttrs
	Download       string
	Href           string
	Hreflang       string
	Ping           string
	Referrerpolicy string
	Rel            string
	Target         string
}

type AArg interface {
	applyA(*AAttrs, *[]Component)
}

func defaultAAttrs() *AAttrs {
	return &AAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func A(args ...AArg) Node {
	a := defaultAAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyA(a, &kids)
	}
	return Node{Tag: "a", Attrs: a, Kids: kids}
}

func (g Global) applyA(a *AAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o DownloadOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Download = o.v
}
func (o HrefOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o PingOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Ping = o.v
}
func (o ReferrerpolicyOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o RelOpt) applyA(a *AAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}
func (o TargetOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Target = o.v
}

func (a *AAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Download != "" {
		Attr(sb, "download", a.Download)
	}
	if a.Href != "" {
		Attr(sb, "href", a.Href)
	}
	if a.Hreflang != "" {
		Attr(sb, "hreflang", a.Hreflang)
	}
	if a.Ping != "" {
		Attr(sb, "ping", a.Ping)
	}
	if a.Referrerpolicy != "" {
		Attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
	if a.Rel != "" {
		Attr(sb, "rel", a.Rel)
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
}
