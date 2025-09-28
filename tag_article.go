package html

import "strings"

type ArticleAttrs struct {
	Global GlobalAttrs
}

type ArticleArg interface {
	ApplyArticle(*ArticleAttrs, *[]Component)
}

func defaultArticleAttrs() *ArticleAttrs {
	return &ArticleAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Article(args ...ArticleArg) Node {
	a := defaultArticleAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyArticle(a, &kids)
	}
	return Node{Tag: "article", Attrs: a, Kids: kids}
}

func (g Global) ApplyArticle(a *ArticleAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *ArticleAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
