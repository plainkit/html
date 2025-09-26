package html

import "strings"

type ScriptAttrs struct {
	Global         GlobalAttrs
	Async          bool
	Attributionsrc string
	Blocking       string
	Crossorigin    string
	Defer          bool
	Fetchpriority  string
	Integrity      string
	Nomodule       string
	Referrerpolicy string
	Src            string
	Type           string
}

type ScriptArg interface {
	applyScript(*ScriptAttrs, *[]Component)
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
		ar.applyScript(a, &kids)
	}
	return Node{Tag: "script", Attrs: a, Kids: kids}
}

func (g Global) applyScript(a *ScriptAttrs, _ *[]Component) {
	g.do(&a.Global)
}

func (o AsyncOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Async = true
}
func (o AttributionsrcOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Attributionsrc = o.v
}
func (o BlockingOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Blocking = o.v
}
func (o CrossoriginOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Crossorigin = o.v
}
func (o DeferOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Defer = true
}
func (o FetchpriorityOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Fetchpriority = o.v
}
func (o IntegrityOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Integrity = o.v
}
func (o NomoduleOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Nomodule = o.v
}
func (o ReferrerpolicyOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Referrerpolicy = o.v
}
func (o SrcOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Src = o.v
}
func (o TypeOpt) applyScript(a *ScriptAttrs, _ *[]Component) {
	a.Type = o.v
}

func (a *ScriptAttrs) writeAttrs(sb *strings.Builder) {
	WriteGlobal(sb, &a.Global)
	if a.Async {
		BoolAttr(sb, "async")
	}
	if a.Attributionsrc != "" {
		Attr(sb, "attributionsrc", a.Attributionsrc)
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
	if a.Nomodule != "" {
		Attr(sb, "nomodule", a.Nomodule)
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
