package html

import "strings"

type RubyAttrs struct {
	Global GlobalAttrs
}

type RubyArg interface {
	ApplyRuby(*RubyAttrs, *[]Component)
}

func defaultRubyAttrs() *RubyAttrs {
	return &RubyAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Ruby(args ...RubyArg) Node {
	a := defaultRubyAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyRuby(a, &kids)
	}

	return Node{Tag: "ruby", Attrs: a, Kids: kids}
}

func (g Global) ApplyRuby(a *RubyAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *RubyAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
