# Plain

Type‑safe HTML for Go. Server‑first, hypermedia by default, htmx as an optional extension. Compile‑time validation, zero hydration, no client runtime.

## What is Plain?

Plain generates HTML using pure Go functions instead of template engines. Each element is a typed function with compile‑time validation and IDE autocomplete. You build pages from links and forms, keep state on the server, and add interactivity progressively (e.g., with htmx) — no SSR/hydration/SPA machinery.

## Why

**Template engines are runtime**: Most Go HTML solutions parse templates at runtime, introduce string-based errors, and lack compile-time guarantees about HTML structure.

**Struct-based builders are verbose**: Existing HTML builders using structs require repetitive field assignments and don't provide intuitive composition patterns.

**Missing type safety**: HTML attributes and structure errors only surface at runtime or in browsers, not during development.

Plain solves these problems by providing compile-time HTML generation with function composition that feels natural in Go.

## How it works

### Function-based composition

```go
div := Div(
    Class("container"),
    H1(T("Hello"), Class("title")),
    P(T("World"), Class("content")),
)
```

### Type-safe attributes per element

Each HTML element has its own option types. Input elements only accept input-specific attributes:

```go
// This compiles and works
input := Input(
    InputType("email"),
    InputName("email"),
    Required(),
    Placeholder("Enter email"),
)

// This fails at compile time - Href is not valid for Input
input := Input(Href("/invalid")) // Compile error
```

### Zero runtime overhead

HTML generation happens through method dispatch resolved at compile time. No reflection, no runtime parsing:

```go
component := Div(Class("test"), T("Hello"))
html := Render(component) // Pure string building
```

## What you get

### IDE support

- **Autocomplete**: Functions and options show up in IDE completion
- **Go to definition**: Jump directly to tag implementations
- **Refactoring**: Rename functions across your entire codebase
- **Type checking**: Invalid attribute combinations fail at compile time

### Modular architecture

Each HTML element lives in its own file (`tag_div.go`, `tag_input.go`, etc.). This makes the codebase:

- Easy to understand and contribute to
- Simple to extend with new elements
- Clear about what's supported

### Testing integration

Components are just Go values. Test them like any other Go code:

```go
func TestButton(t *testing.T) {
    btn := Button(
        ButtonType("submit"),
        T("Click me"),
        Class("btn"),
    )

    html := Render(btn)
    if !strings.Contains(html, `type="submit"`) {
        t.Error("Missing type attribute")
    }
}
```

### Composition patterns

Build reusable components by composing smaller ones:

```go
func Card(title, content string) Node {
    return Div(
        Class("card"),
        H2(T(title), Class("card-title")),
        P(T(content), Class("card-content")),
    )
}

func Page() Node {
    return Div(
        Class("container"),
        Card("Welcome", "Get started with Plain"),
        Card("Features", "Type-safe HTML in Go"),
    )
}
```

## Installation

```bash
go get github.com/plainkit/html
```

## Usage

### Basic example

```go
package main

import (
    "fmt"
    . "github.com/plainkit/html"
)

func main() {
    page := Html(
        Lang("en"),
        Head(
            HeadTitle(T("My Page")),
            Meta(Charset("UTF-8")),
			HeadStyle(T(".intro { color: blue; }")),
        ),
        Body(
            H1(T("Hello, World!")),
            P(T("Built with Plain"), Class("intro")),
        ),
    )

    fmt.Println("<!DOCTYPE html>")
    fmt.Println(Render(page))
}
```

### Working with forms

```go
loginForm := Form(
    Action("/login"),
    Method("POST"),
    Div(
        Label(For("email"), T("Email")),
        Input(
            InputType("email"),
            InputName("email"),
            Id("email"),
            Required(),
        ),
    ),
    Div(
        Label(For("password"), T("Password")),
        Input(
            InputType("password"),
            InputName("password"),
            Id("password"),
            Required(),
        ),
    ),
    Button(
        ButtonType("submit"),
        T("Login"),
    ),
)
```

### HTTP handlers

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    page := Html(
        Lang("en"),
        Head(HeadTitle(T("Home"))),
        Body(
            H1(T("Welcome")),
            P(T("This page was built with Plain")),
        ),
    )

    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<!DOCTYPE html>\n")
    fmt.Fprint(w, Render(page))
}
```

## Architecture

### Core types

- **Component**: Interface implemented by all HTML elements
- **Node**: Struct representing an HTML element with tag, attributes, and children
- **TextNode**: Represents plain text content
- **Global**: Option type that works with any HTML element (class, id, etc.)

### Tag-specific options

Each HTML element has dedicated option types:

- `InputType()`, `InputName()`, `Required()` for `Input()`
- `Href()`, `Target()`, `Rel()` for `A()`
- `ButtonType()`, `Disabled()` for `Button()`
- And so on for all HTML5 elements

### File organization

```
├── core_node.go          # Component interface, Node struct, Render()
├── core_global.go        # GlobalAttrs struct, attribute helpers
├── core_options.go       # Global option constructors (Class, Id, etc.)
├── core_content.go       # Text() and Child() helpers
├── tag_div.go           # Div component and options
├── tag_input.go         # Input component and options
├── tag_form.go          # Form, Input, Textarea, Button, etc.
├── tag_semantic.go      # Header, Nav, Main, Section, etc.
└── ...                  # One file per logical group of elements
```

## Contributing

The codebase is designed for easy contribution:

1. **Add a new HTML element**: Create `tag_newelem.go` following existing patterns
2. **Add element-specific attributes**: Define option types and `apply*` methods
3. **Test**: Add examples to `sample/` directory
4. **Document**: Update this README with usage examples

Each tag file follows the same pattern:

- Attrs struct with element-specific fields
- Arg interface for type safety
- Constructor function accepting variadic args
- Option types and constructors
- Apply methods connecting options to attributes
- writeAttrs method for HTML generation

## License

MIT
