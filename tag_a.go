package html

import "strings"

type AAttrs struct {
	Global              GlobalAttrs
	Attributionsourceid string
	Attributionsrc      string
	Charset             string
	Coords              string
	Download            string
	Href                string
	Hreflang            string
	Hreftranslate       string
	ImplicitNoopener    string
	Name                string
	Ping                string
	Referrerpolicy      string
	Rel                 string
	Rev                 string
	Shape               string
	Target              string
	TextFragments       string
	Type                string
}

type AArg interface {
	applyA(*AAttrs, *[]Component)
}

func defaultAAttrs() *AAttrs {
	return &AAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func A(args ...AArg) Node {
	a := defaultAAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyA(a, &kids)
	}
	return Node{Tag: "a", Attrs: a, Kids: kids}
}

func (g Global) applyA(a *AAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AttributionsourceidOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Attributionsourceid = o.v
}
func (o AttributionsrcOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Attributionsrc = o.v
}
func (o CharsetOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Charset = o.v
}
func (o CoordsOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Coords = o.v
}
func (o DownloadOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Download = o.v
}
func (o HrefOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o HreftranslateOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Hreftranslate = o.v
}
func (o ImplicitNoopenerOpt) applyA(a *AAttrs, _ *[]Component) {
	a.ImplicitNoopener = o.v
}
func (o NameOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Name = o.v
}
func (o PingOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Ping = o.v
}
func (o ReferrerpolicyOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o RelOpt) applyA(a *AAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}
func (o RevOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Rev = o.v
}
func (o ShapeOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Shape = o.v
}
func (o TargetOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Target = o.v
}
func (o TextFragmentsOpt) applyA(a *AAttrs, _ *[]Component) {
	a.TextFragments = o.v
}
func (o TypeOpt) applyA(a *AAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *AAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Attributionsourceid != "" {
		Attr(sb, "attributionsourceid", a.Attributionsourceid)
	}
	if a.Attributionsrc != "" {
		Attr(sb, "attributionsrc", a.Attributionsrc)
	}
	if a.Charset != "" {
		Attr(sb, "charset", a.Charset)
	}
	if a.Coords != "" {
		Attr(sb, "coords", a.Coords)
	}
	if a.Download != "" {
		Attr(sb, "download", a.Download)
	}
	if a.Href != "" {
		Attr(sb, "href", a.Href)
	}
	if a.Hreflang != "" {
		Attr(sb, "hreflang", a.Hreflang)
	}
	if a.Hreftranslate != "" {
		Attr(sb, "hreftranslate", a.Hreftranslate)
	}
	if a.ImplicitNoopener != "" {
		Attr(sb, "implicit_noopener", a.ImplicitNoopener)
	}
	if a.Name != "" {
		Attr(sb, "name", a.Name)
	}
	if a.Ping != "" {
		Attr(sb, "ping", a.Ping)
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
	if a.Shape != "" {
		Attr(sb, "shape", a.Shape)
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
	if a.TextFragments != "" {
		Attr(sb, "text_fragments", a.TextFragments)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
