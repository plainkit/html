package html

import "strings"

type LinkAttrs struct {
	Global         GlobalAttrs
	As             string
	Blocking       string
	Color          string
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
	Sizes          string
	Type           string
}

type LinkArg interface {
	ApplyLink(*LinkAttrs, *[]Component)
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
		ar.ApplyLink(a, &kids)
	}

	return Node{Tag: "link", Attrs: a, Kids: kids, Void: true}
}

func (g Global) ApplyLink(a *LinkAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AsOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.As = o.v
}
func (o BlockingOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Blocking = o.v
}
func (o ColorOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Color = o.v
}
func (o CrossoriginOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DisabledOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Disabled = true
}
func (o FetchpriorityOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o HrefOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o ImagesizesOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Imagesizes = o.v
}
func (o ImagesrcsetOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Imagesrcset = o.v
}
func (o IntegrityOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Integrity = o.v
}
func (o MediaOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Media = o.v
}
func (o ReferrerpolicyOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o RelOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}
func (o SizesOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Sizes = o.v
}
func (o TypeOpt) ApplyLink(a *LinkAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *LinkAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)

	if a.As != "" {
		Attr(sb, "as", a.As)
	}

	if a.Blocking != "" {
		Attr(sb, "blocking", a.Blocking)
	}

	if a.Color != "" {
		Attr(sb, "color", a.Color)
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

	if a.Sizes != "" {
		Attr(sb, "sizes", a.Sizes)
	}

	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
