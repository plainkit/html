package html

import "strings"

type HtmlAttrs struct {
	Global GlobalAttrs
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
	g.do(&a.Global)
}

func (a *HtmlAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
