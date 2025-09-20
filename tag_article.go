package html

import "strings"

// Article
type ArticleAttrs struct {
	Global GlobalAttrs
}

type ArticleArg interface {
	applyArticle(*ArticleAttrs, *[]Component)
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
		ar.applyArticle(a, &kids)
	}
	return Node{Tag: "article", Attrs: a, Kids: kids}
}

func (g Global) applyArticle(a *ArticleAttrs, _ *[]Component) { g.do(&a.Global) }
func (o TxtOpt) applyArticle(_ *ArticleAttrs, kids *[]Component) {
	*kids = append(*kids, TextNode(o.s))
}
func (o ChildOpt) applyArticle(_ *ArticleAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *ArticleAttrs) writeAttrs(sb *strings.Builder)             { writeGlobal(sb, &a.Global) }
