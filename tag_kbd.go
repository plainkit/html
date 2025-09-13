package blox

import "strings"

// Kbd
type KbdAttrs struct {
	Global GlobalAttrs
}

type KbdArg interface {
	applyKbd(*KbdAttrs, *[]Component)
}

func defaultKbdAttrs() *KbdAttrs {
	return &KbdAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Kbd(args ...KbdArg) Component {
	a := defaultKbdAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyKbd(a, &kids)
	}
	return Node{Tag: "kbd", Attrs: a, Kids: kids}
}

func (g Global) applyKbd(a *KbdAttrs, _ *[]Component)      { g.do(&a.Global) }
func (o TxtOpt) applyKbd(_ *KbdAttrs, kids *[]Component)   { *kids = append(*kids, TextNode(o.s)) }
func (o ChildOpt) applyKbd(_ *KbdAttrs, kids *[]Component) { *kids = append(*kids, o.c) }
func (a *KbdAttrs) writeAttrs(sb *strings.Builder)         { writeGlobal(sb, &a.Global) }
