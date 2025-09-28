package generator

const tagsTemplate = `package html

import "strings"

type {{.StructName}} struct {
	Global GlobalAttrs{{range .Attributes}}
	{{.Field}} {{.GoType}}{{end}}
}

type {{.ArgInterface}} interface {
	Apply{{.Title}}(*{{.StructName}}, *[]Component)
}

func default{{.StructName}}() *{{.StructName}} {
	return &{{.StructName}}{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func {{.Title}}(args ...{{.ArgInterface}}) Node {
	a := default{{.StructName}}()
	var kids []Component
	for _, ar := range args {
		ar.Apply{{.Title}}(a, &kids)
}
	{{if .Void}}return Node{Tag: "{{.Name}}", Attrs: a, Kids: kids, Void: true}{{else}}return Node{Tag: "{{.Name}}", Attrs: a, Kids: kids}{{end}}
}

func (g Global) Apply{{.Title}}(a *{{.StructName}}, _ *[]Component) {
	g.Do(&a.Global)
}
{{range .Attributes}}{{if ne .Attr "data"}}
func (o {{.Field}}Opt) Apply{{$.Title}}(a *{{$.StructName}}, _ *[]Component) {
	{{if eq .Type "bool"}}a.{{.Field}} = true{{else if eq .Attr "rel"}}if a.{{.Field}} == "" {
		a.{{.Field}} = o.v
	} else {
		a.{{.Field}} += " " + o.v
	}{{else}}a.{{.Field}} = o.v{{end}}
}{{end}}{{end}}

func (a *{{.StructName}}) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global){{range .Attributes}}{{if ne .Attr "data"}}
	{{if eq .Type "bool"}}if a.{{.Field}} {
		BoolAttr(sb, "{{.Attr}}")
	}{{else}}if a.{{.Field}} != "" {
		Attr(sb, "{{.Attr}}", a.{{.Field}})
	}{{end}}{{end}}{{end}}
}
`
