package blox

import "strings"

// HR (void)
type HrAttrs struct {
	Global GlobalAttrs
}

type HrArg interface {
	applyHr(*HrAttrs)
}

func defaultHrAttrs() *HrAttrs {
	return &HrAttrs{
		Global: GlobalAttrs{
			Style:  map[string]string{},
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Hr(args ...HrArg) Node {
	a := defaultHrAttrs()
	for _, ar := range args {
		ar.applyHr(a)
	}
	return Node{Tag: "hr", Attrs: a, Void: true}
}

func (g Global) applyHr(a *HrAttrs)               { g.do(&a.Global) }
func (a *HrAttrs) writeAttrs(sb *strings.Builder) { writeGlobal(sb, &a.Global) }
