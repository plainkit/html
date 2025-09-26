package generator

const attributesTemplate = `package html


{{range .Attributes}}{{if ne .Attr "data"}}{{if eq .Type "bool"}}type {{.Field}}Opt struct{}

func A{{.Field}}() {{.Field}}Opt {
	return {{.Field}}Opt{}
}

{{else}}type {{.Field}}Opt struct {
	v {{.GoType}}
}

func A{{.Field}}(v {{.GoType}}) {{.Field}}Opt {
	return {{.Field}}Opt{v}
}

{{end}}{{end}}{{end}}`
