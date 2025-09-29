package html

import "strings"

type KbdAttrs struct {
	Global GlobalAttrs
}

type KbdArg interface {
	ApplyKbd(*KbdAttrs, *[]Component)
}

func defaultKbdAttrs() *KbdAttrs {
	return &KbdAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Kbd(args ...KbdArg) Node {
	a := defaultKbdAttrs()

	var kids []Component
	for _, ar := range args {
		ar.ApplyKbd(a, &kids)
	}

	return Node{Tag: "kbd", Attrs: a, Kids: kids}
}

func (g Global) ApplyKbd(a *KbdAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (a *KbdAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
