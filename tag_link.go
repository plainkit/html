package html

import "strings"

type LinkAttrs struct {
	Global         GlobalAttrs
	As             string
	Blocking       string
	Charset        string
	Crossorigin    string
	Disabled       bool
	Fetchpriority  string
	Href           string
	Hreflang       string
	Imagesizes     string
	Imagesrcset    string
	Integrity      string
	Media          string
	Referrerpolicy string
	Rel            string
	Rev            string
	Sizes          string
	Target         string
	Type           string
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
func (o BlockingOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Blocking = o.v
}
func (o CharsetOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Charset = o.v
}
func (o CrossoriginOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DisabledOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FetchpriorityOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o HrefOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o ImagesizesOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Imagesizes = o.v
}
func (o ImagesrcsetOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Imagesrcset = o.v
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
func (o RevOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Rev = o.v
}
func (o SizesOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Sizes = o.v
}
func (o TargetOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Target = o.v
}
func (o TypeOpt) applyLink(a *LinkAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *LinkAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.As != "" {
		Attr(sb, "as", a.As)
	}
	if a.Blocking != "" {
		Attr(sb, "blocking", a.Blocking)
	}
	if a.Charset != "" {
		Attr(sb, "charset", a.Charset)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Disabled {
		BoolAttr(sb, "disabled")
	}
	if a.Fetchpriority != "" {
		Attr(sb, "fetchpriority", a.Fetchpriority)
	}
	if a.Href != "" {
		Attr(sb, "href", a.Href)
	}
	if a.Hreflang != "" {
		Attr(sb, "hreflang", a.Hreflang)
	}
	if a.Imagesizes != "" {
		Attr(sb, "imagesizes", a.Imagesizes)
	}
	if a.Imagesrcset != "" {
		Attr(sb, "imagesrcset", a.Imagesrcset)
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
	if a.Rev != "" {
		Attr(sb, "rev", a.Rev)
	}
	if a.Sizes != "" {
		Attr(sb, "sizes", a.Sizes)
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
