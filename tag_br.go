package blox

import "strings"

// BR (void)
type BrAttrs struct {
	Global GlobalAttrs
}

type BrArg interface {
	applyBr(*BrAttrs)
}

func defaultBrAttrs() *BrAttrs {
	return &BrAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Br(args ...BrArg) Component {
	a := defaultBrAttrs()
	for _, ar := range args {
		ar.applyBr(a)
	}
	return Node{Tag: "br", Attrs: a, Void: true}
}

func (g Global) applyBr(a *BrAttrs)               { g.do(&a.Global) }
func (a *BrAttrs) writeAttrs(sb *strings.Builder) { writeGlobal(sb, &a.Global) }
