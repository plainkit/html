package html

import "strings"

type HtmlAttrs struct {
	Global   GlobalAttrs
	Manifest string
	Version  string
	Xmlns    string
}

type HtmlArg interface {
	applyHtml(*HtmlAttrs, *[]Component)
}

func defaultHtmlAttrs() *HtmlAttrs {
	return &HtmlAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Html(args ...HtmlArg) Node {
	a := defaultHtmlAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyHtml(a, &kids)
	}
	return Node{Tag: "html", Attrs: a, Kids: kids}
}

// Tag-specific options
type ManifestOpt struct{ v string }
type VersionOpt struct{ v string }
type XmlnsOpt struct{ v string }

func Manifest(v string) ManifestOpt { return ManifestOpt{v} }
func Version(v string) VersionOpt   { return VersionOpt{v} }
func Xmlns(v string) XmlnsOpt       { return XmlnsOpt{v} }

// Global option glue
func (g Global) applyHtml(a *HtmlAttrs, _ *[]Component) {
	g.do(&a.Global)
}

// Content option glue
func (o TxtOpt) applyHtml(_ *HtmlAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}

func (o ChildOpt) applyHtml(_ *HtmlAttrs, kids *[]Component) {
	*kids = append(*kids, o.c)
}

// Tag-specific option glue
func (o ManifestOpt) applyHtml(a *HtmlAttrs, _ *[]Component) { a.Manifest = o.v }
func (o VersionOpt) applyHtml(a *HtmlAttrs, _ *[]Component)  { a.Version = o.v }
func (o XmlnsOpt) applyHtml(a *HtmlAttrs, _ *[]Component)    { a.Xmlns = o.v }

// Attrs writer implementation
func (a *HtmlAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Manifest != "" {
		attr(sb, "manifest", a.Manifest)
	}
	if a.Version != "" {
		attr(sb, "version", a.Version)
	}
	if a.Xmlns != "" {
		attr(sb, "xmlns", a.Xmlns)
	}
}
