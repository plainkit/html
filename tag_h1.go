package html

import "strings"

type H1Attrs struct {
	Global                             GlobalAttrs
	NoUaStylesInArticleAsideNavSection string
}

type H1Arg interface {
	applyH1(*H1Attrs, *[]Component)
}

func defaultH1Attrs() *H1Attrs {
	return &H1Attrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func H1(args ...H1Arg) Node {
	a := defaultH1Attrs()
	var kids []Component
	for _, ar := range args {
		ar.applyH1(a, &kids)
	}
	return Node{Tag: "h1", Attrs: a, Kids: kids}
}

func (g Global) applyH1(a *H1Attrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o NoUaStylesInArticleAsideNavSectionOpt) applyH1(a *H1Attrs, _ *[]Component) {
	a.NoUaStylesInArticleAsideNavSection = o.v
}

func (a *H1Attrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.NoUaStylesInArticleAsideNavSection != "" {
		Attr(sb, "no_ua_styles_in_article_aside_nav_section", a.NoUaStylesInArticleAsideNavSection)
	}
}
