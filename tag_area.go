package html

import "strings"

type AreaAttrs struct {
	Global         GlobalAttrs
	Alt            string
	Coords         string
	Download       string
	Href           string
	Hreflang       string
	Nohref         bool
	Ping           string
	Referrerpolicy string
	Rel            string
	Shape          string
	Target         string
	Type           string
}

type AreaArg interface {
	applyArea(*AreaAttrs, *[]Component)
}

func defaultAreaAttrs() *AreaAttrs {
	return &AreaAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Area(args ...AreaArg) Node {
	a := defaultAreaAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyArea(a, &kids)
	}
	return Node{Tag: "area", Attrs: a, Kids: kids, Void: true}
}

func (g Global) applyArea(a *AreaAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AltOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Alt = o.v
}
func (o CoordsOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Coords = o.v
}
func (o DownloadOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Download = o.v
}
func (o HrefOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Href = o.v
}
func (o HreflangOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Hreflang = o.v
}
func (o NohrefOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Nohref = true
}
func (o PingOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Ping = o.v
}
func (o ReferrerpolicyOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o RelOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	if a.Rel == "" {
		a.Rel = o.v
	} else {
		a.Rel += " " + o.v
	}
}
func (o ShapeOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Shape = o.v
}
func (o TargetOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Target = o.v
}
func (o TypeOpt) applyArea(a *AreaAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *AreaAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Alt != "" {
		Attr(sb, "alt", a.Alt)
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
	if a.Nohref {
		BoolAttr(sb, "nohref")
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
	if a.Shape != "" {
		Attr(sb, "shape", a.Shape)
	}
	if a.Target != "" {
		Attr(sb, "target", a.Target)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
