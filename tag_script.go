package html

import "strings"

// Script
type ScriptAttrs struct {
	Global GlobalAttrs
	Src    string
	Type   string
	Async  bool
	Defer  bool
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

type ScriptComponent Node

func (script ScriptComponent) render(sb *strings.Builder) {
	Node(script).render(sb)
}

func Script(args ...ScriptArg) ScriptComponent {
	a := defaultScriptAttrs()
	var kids []Component
	for _, ar := range args {
		ar.applyScript(a, &kids)
	}
	return ScriptComponent{Tag: "script", Attrs: a, Kids: kids}
}

// Script-specific options
type AsyncOpt struct{}
type DeferOpt struct{}
type ScriptSrcOpt struct{ v string }
type ScriptTypeOpt struct{ v string }

func Async() AsyncOpt                   { return AsyncOpt{} }
func Defer() DeferOpt                   { return DeferOpt{} }
func ScriptSrc(v string) ScriptSrcOpt   { return ScriptSrcOpt{v} }
func ScriptType(v string) ScriptTypeOpt { return ScriptTypeOpt{v} }

func (g Global) applyScript(a *ScriptAttrs, _ *[]Component)    { g.do(&a.Global) }
func (o TxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) { *kids = append(*kids, TextNode(o.s)) }
func (o UnsafeTxtOpt) applyScript(_ *ScriptAttrs, kids *[]Component) {
	*kids = append(*kids, UnsafeTextNode(o.s))
}
func (o ChildOpt) applyScript(_ *ScriptAttrs, kids *[]Component)   { *kids = append(*kids, o.c) }
func (o AsyncOpt) applyScript(a *ScriptAttrs, _ *[]Component)      { a.Async = true }
func (o DeferOpt) applyScript(a *ScriptAttrs, _ *[]Component)      { a.Defer = true }
func (o ScriptSrcOpt) applyScript(a *ScriptAttrs, _ *[]Component)  { a.Src = o.v }
func (o ScriptTypeOpt) applyScript(a *ScriptAttrs, _ *[]Component) { a.Type = o.v }

func (script ScriptComponent) applyHead(_ *HeadAttrs, kids *[]Component) {
	*kids = append(*kids, script)
}

func (script ScriptComponent) applyBody(_ *BodyAttrs, kids *[]Component) {
	*kids = append(*kids, script)
}

func (a *ScriptAttrs) writeAttrs(sb *strings.Builder) {
	writeGlobal(sb, &a.Global)
	if a.Src != "" {
		attr(sb, "src", a.Src)
	}
	if a.Type != "" {
		attr(sb, "type", a.Type)
	}
	if a.Async {
		boolAttr(sb, "async")
	}
	if a.Defer {
		boolAttr(sb, "defer")
	}
}
