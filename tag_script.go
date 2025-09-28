package html

import "strings"

type ScriptAttrs struct {
	Global         GlobalAttrs
	Async          bool
	Blocking       string
	Crossorigin    string
	Defer          bool
	Fetchpriority  string
	Integrity      string
	Nomodule       bool
	Referrerpolicy string
	Src            string
	Type           string
}

type ScriptArg interface {
	ApplyScript(*ScriptAttrs, *[]Component)
}

func defaultScriptAttrs() *ScriptAttrs {
	return &ScriptAttrs{
		Global: GlobalAttrs{
			Style:  "",
			Aria:   map[string]string{},
			Data:   map[string]string{},
			Events: map[string]string{},
		},
	}
}

func Script(args ...ScriptArg) Node {
	a := defaultScriptAttrs()
	var kids []Component
	for _, ar := range args {
		ar.ApplyScript(a, &kids)
	}
	return Node{Tag: "script", Attrs: a, Kids: kids}
}

func (g Global) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	g.Do(&a.Global)
}

func (o AsyncOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Async = true
}
func (o BlockingOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Blocking = o.v
}
func (o CrossoriginOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DeferOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Defer = true
}
func (o FetchpriorityOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o IntegrityOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Integrity = o.v
}
func (o NomoduleOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Nomodule = true
}
func (o ReferrerpolicyOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o SrcOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o TypeOpt) ApplyScript(a *ScriptAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *ScriptAttrs) WriteAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Async {
		BoolAttr(sb, "async")
	}
	if a.Blocking != "" {
		Attr(sb, "blocking", a.Blocking)
	}
	if a.Crossorigin != "" {
		Attr(sb, "crossorigin", a.Crossorigin)
	}
	if a.Defer {
		BoolAttr(sb, "defer")
	}
	if a.Fetchpriority != "" {
		Attr(sb, "fetchpriority", a.Fetchpriority)
	}
	if a.Integrity != "" {
		Attr(sb, "integrity", a.Integrity)
	}
	if a.Nomodule {
		BoolAttr(sb, "nomodule")
	}
	if a.Referrerpolicy != "" {
		Attr(sb, "referrerpolicy", a.Referrerpolicy)
	}
	if a.Src != "" {
		Attr(sb, "src", a.Src)
	}
	if a.Type != "" {
		Attr(sb, "type", a.Type)
	}
}
