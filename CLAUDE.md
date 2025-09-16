# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

description: Guidelines for creating Plain idiomatic HTML components
globs: "\*.go"
alwaysApply: true

---

## Project Overview

Plain is a function-based HTML component library for Go that generates HTML at compile time with zero runtime overhead. It provides type-safe HTML construction through function composition instead of template engines or struct-based builders.

## Development Commands

### Core Library Development

- **Build**: `go build` (builds the main library)
- **Test**: `go test ./...` (run all tests)
- **Run sample**: `go run sample/main.go` (if sample exists)
- **Format**: `go fmt ./...`
- **Vet**: `go vet ./...`

## Core Architecture

### Component System

- **Component**: Interface with `render(*strings.Builder)` method
- **Node**: Core struct representing HTML elements (tag + attributes + children)
- **TextNode**: Represents escaped text content
- **Global**: Universal options that work with any HTML element

### File Organization Pattern

```
core_*.go           # Core interfaces and rendering system (core_node.go, core_options.go, core_content.go)
tag_*.go           # HTML element implementations (tag_div.go, tag_input.go, etc.)
assets.go          # Asset collection and management system
go.mod             # Module definition
```

### HTML Element Architecture

Each `tag_*.go` file follows this pattern:

1. **Attrs struct**: Element-specific attributes (e.g., `DivAttrs`, `InputAttrs`)
2. **Arg interface**: Type constraint for element options (e.g., `DivArg`, `InputArg`)
3. **Constructor function**: Main element function accepting variadic args
4. **Option types**: Specific attribute setters (e.g., `InputType()`, `Required()`)
5. **Apply methods**: Connect options to attributes via interface satisfaction
6. **writeAttrs method**: Generates HTML attribute strings

### Asset System

The library includes an asset collection system that allows components to declare CSS and JavaScript dependencies:

- **Node.AssetCSS**: CSS content to be collected
- **Node.AssetJS**: JavaScript content to be collected
- **Node.AssetName**: Name for asset deduplication
- **WithAssets()**: Method to attach assets to components

## Type Safety System

### Compile-Time Guarantees

- Each HTML element only accepts valid attributes for that element
- Invalid combinations like `Input(Href("/invalid"))` fail at compile time
- Children and text content are type-safe through the Component interface

### Attribute Validation

- Global attributes (class, id, etc.) work with any element via `Global` type
- Element-specific attributes are constrained by their respective `Arg` interfaces
- Method dispatch resolves at compile time with zero runtime overhead

## Content and Text Handling

### Text Content

- **TextNode**: Automatically HTML-escapes content for security
- **UnsafeTextNode**: Raw HTML content (use with caution)
- **T()**: Convenience function for creating TextNode
- **Child()**: Function for adding component children

### Component Composition

- All HTML elements can accept other components as children
- Node implements apply methods for all major HTML elements
- Direct component passing without wrapping (e.g., `Div(myComponent)` works)

## Adding New Elements

### Core HTML Elements

1. Create `tag_newelement.go` following the established pattern
2. Define element-specific `Attrs` struct and `Arg` interface
3. Implement constructor function with variadic arguments
4. Add element-specific option constructors and apply methods
5. Implement `writeAttrs()` method for HTML generation

## Global Attribute System

### Universal Attributes

The `Global` type provides attributes that work with any HTML element:

- **Class()**: CSS class names
- **Id()**: Element ID
- **Data()**: Data attributes (data-\*)
- **Aria()**: ARIA accessibility attributes
- **Style()**: Inline CSS styles
- **On()**: Event handlers
- **Custom()**: Custom attributes (like hx-_, x-_, etc.)

### Content Functions

- **T()**: Create escaped text content
- **Child()**: Add component children
- **C()**: Shorthand for Child()

## HTML Element Coverage

The library provides comprehensive HTML5 element support:

- **Document structure**: Html, Head, Body, Title, Meta, Link, Style, Script
- **Sectioning**: Header, Nav, Main, Section, Article, Aside, Footer
- **Headings**: H1-H6
- **Text content**: P, Div, Span, Pre, Blockquote, List elements (Ul, Ol, Li)
- **Text semantics**: A, Strong, Em, Small, S, Cite, Q, Dfn, Abbr, Time, Code, Var, Samp, Kbd, Sub, Sup, I, B, U, Mark, Del, Ins
- **Forms**: Form, Input, Textarea, Button, Select, Option, Optgroup, Label, Fieldset, Legend, Datalist
- **Tables**: Table, Thead, Tbody, Tfoot, Tr, Th, Td, Caption, Colgroup, Col
- **Media**: Img, Audio, Video, Source, Track, Canvas, Iframe
- **Interactive**: Details, Summary, Dialog
- **SVG**: Svg and common SVG elements (Circle, Rect, Path, etc.)

## Writing Tests

Create test files following Go conventions:

```go
func TestComponent(t *testing.T) {
    component := Div(
        Class("test"),
        T("Hello"),
    )

    html := Render(component)
    expected := `<div class="test">Hello</div>`

    if html != expected {
        t.Errorf("Expected %q, got %q", expected, html)
    }
}
```
