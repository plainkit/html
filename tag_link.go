package html

import "strings"

// Link (void)
type LinkAttrs struct {
	Global      GlobalAttrs
	Href        string
	Rel         string
	Type        string
	Media       string
	Hreflang    string
	Sizes       string
	Crossorigin string
}

type LinkArg interface {
	applyLink(*LinkAttrs)
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

type LinkComponent Node

func (link LinkComponent) render(sb *strings.Builder) {
	Node(link).render(sb)
}

func Link(args ...LinkArg) LinkComponent {
	a := defaultLinkAttrs()
	for _, ar := range args {
		ar.applyLink(a)
	}
	return LinkComponent{Tag: "link", Attrs: a, Void: true}
}

// Link-specific options
type LinkHrefOpt struct{ v string }
type LinkRelOpt struct{ v string }
type LinkTypeOpt struct{ v string }
type HreflangOpt struct{ v string }
type SizesOpt struct{ v string }
type CrossoriginOpt struct{ v string }
type LinkMediaOpt struct{ v string }

func LinkHref(v string) LinkHrefOpt       { return LinkHrefOpt{v} }
func LinkRel(v string) LinkRelOpt         { return LinkRelOpt{v} }
func LinkType(v string) LinkTypeOpt       { return LinkTypeOpt{v} }
func Hreflang(v string) HreflangOpt       { return HreflangOpt{v} }
func Sizes(v string) SizesOpt             { return SizesOpt{v} }
func Crossorigin(v string) CrossoriginOpt { return CrossoriginOpt{v} }
func LinkMedia(v string) LinkMediaOpt     { return LinkMediaOpt{v} }

func (g Global) applyLink(a *LinkAttrs)         { g.do(&a.Global) }
func (o LinkHrefOpt) applyLink(a *LinkAttrs)    { a.Href = o.v }
func (o LinkRelOpt) applyLink(a *LinkAttrs)     { a.Rel = o.v }
func (o LinkTypeOpt) applyLink(a *LinkAttrs)    { a.Type = o.v }
func (o HreflangOpt) applyLink(a *LinkAttrs)    { a.Hreflang = o.v }
func (o SizesOpt) applyLink(a *LinkAttrs)       { a.Sizes = o.v }
func (o CrossoriginOpt) applyLink(a *LinkAttrs) { a.Crossorigin = o.v }
func (o LinkMediaOpt) applyLink(a *LinkAttrs)   { a.Media = o.v }

// Compile-time type safety: Link can be added to Head
func (link LinkComponent) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, link)
}

func (a *LinkAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Href != "" {
		attr(sb, "href", a.Href)
	}
	if a.Rel != "" {
		attr(sb, "rel", a.Rel)
	}
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Media != "" {
		attr(sb, "media", a.Media)
	}
	if a.Hreflang != "" {
		attr(sb, "hreflang", a.Hreflang)
	}
	if a.Sizes != "" {
		attr(sb, "sizes", a.Sizes)
	}
	if a.Crossorigin != "" {
		attr(sb, "crossorigin", a.Crossorigin)
	}
}
