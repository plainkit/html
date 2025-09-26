package html

import "strings"

type FileAttrs struct {
	Global GlobalAttrs
}

type FileArg interface {
	applyFile(*FileAttrs, *[]Component)
}

func defaultFileAttrs() *FileAttrs {
	return &FileAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func File(args ...FileArg) Node {
	a := defaultFileAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyFile(a, &kids)
	}
	return Node{Tag: "file", Attrs: a, Kids: kids}
}

func (g Global) applyFile(a *FileAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (a *FileAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
}
