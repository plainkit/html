package html

import "strings"

type LinkAttrs struct {
	Global         GlobalAttrs
	As             string
	Crossorigin    string
	Href           string
	Hreflang       string
	Integrity      string
	Media          string
	Referrerpolicy string
	Rel            string
	Sizes          string
}

type LinkArg interface {
	applyLink(*LinkAttrs, *[]Component)
}

func defaultLinkAttrs() *LinkAttrs {
	return &LinkAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Link(args ...LinkArg) Node {
	a := defaultLinkAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyLink(a, &kids)
	}
	return Node{Tag: "link", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyLink(a *LinkAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AsOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.As = o.v
}
func (o CrossoriginOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o HrefOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o IntegrityOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Integrity = o.v
}
func (o MediaOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o ReferrerpolicyOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o RelOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}
func (o SizesOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Sizes = o.v
}

func (a *LinkAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.As != "" {
		Attr(sb, "as", a.As)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Href != "" {
		Attr(sb, "href", a.Href)
	}
	if a.Hreflang != "" {
		Attr(sb, "hreflang", a.Hreflang)
	}
	if a.Integrity != "" {
		Attr(sb, "integrity", a.Integrity)
	}
	if a.Media != "" {
		Attr(sb, "media", a.Media)
	}
	if a.Referrerpolicy != "" {
		Attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
	if a.Rel != "" {
		Attr(sb, "rel", a.Rel)
	}
	if a.Sizes != "" {
		Attr(sb, "sizes", a.Sizes)
	}
}
