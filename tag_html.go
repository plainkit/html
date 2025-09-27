package html

import "strings"

type HtmlAttrs struct {
	Global   GlobalAttrs
	Manifest string
	Version  string
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

func (g Global) applyHtml(a *HtmlAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o ManifestOpt) applyHtml(a *HtmlAttrs, _ *[]Component) {
	a.Manifest = o.v
}
func (o VersionOpt) applyHtml(a *HtmlAttrs, _ *[]Component) {
	a.Version = o.v
}

func (a *HtmlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Manifest != "" {
		Attr(sb, "manifest", a.Manifest)
	}
	if a.Version != "" {
		Attr(sb, "version", a.Version)
	}
}
