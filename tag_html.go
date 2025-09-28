package html

import "strings"

type HtmlAttrs struct {
	Global GlobalAttrs
}

type HtmlArg interface {
	ApplyHtml(*HtmlAttrs, *[]Component)
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
		ar.ApplyHtml(a, &kids)
	}
	return Node{Tag: "html", Attrs: a, Kids: kids}
}

func (g Global) ApplyHtml(a *HtmlAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *HtmlAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
