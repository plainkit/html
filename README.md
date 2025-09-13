# Blox

A function-based HTML component library for Go that generates HTML at compile time with zero runtime overhead.

## Why

**Template engines are runtime**: Most Go HTML solutions parse templates at runtime, introduce string-based errors, and lack compile-time guarantees about HTML structure.

**Struct-based builders are verbose**: Existing HTML builders using structs require repetitive field assignments and don't provide intuitive composition patterns.

**Missing type safety**: HTML attributes and structure errors only surface at runtime or in browsers, not during development.

Blox solves these problems by providing compile-time HTML generation with function composition that feels natural in Go.

## How it works

### Function-based composition

```go
div := Div(
    Class("container"),
    Child(H1(Text("Hello"), Class("title"))),
    Child(P(Text("World"), Class("content"))),
)
```

### Type-safe attributes per element

Each HTML element has its own option types. Input elements only accept input-specific attributes:

```go
// This compiles
input := Input(
    InputType("email"),
    InputName("email"),
    Required(),
    Placeholder("Enter email"),
)

// This won't compile - Href is not valid for Input
input := Input(Href("/invalid")) // Compile error
```

### Zero runtime overhead

HTML generation happens through method dispatch resolved at compile time. No reflection, no runtime parsing:

```go
component := Div(Class("test"), Text("Hello"))
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
        Text("Click me"),
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
func Card(title, content string) Component {
    return Div(
        Class("card"),
        Child(H2(Text(title), Class("card-title"))),
        Child(P(Text(content), Class("card-content"))),
    )
}

func Page() Component {
    return Div(
        Class("container"),
        Child(Card("Welcome", "Get started with Blox")),
        Child(Card("Features", "Type-safe HTML in Go")),
    )
}
```

## Installation

```bash
go get github.com/bloxui/blox
```

## Usage

### Basic example

```go
package main

import (
    "fmt"
    . "github.com/bloxui/blox"
)

func main() {
    page := Html(
        Lang("en"),
        Child(Head(
            Child(HeadTitle(Text("My Page"))),
            Child(Meta(Charset("UTF-8"))),
        )),
        Child(Body(
            Child(H1(Text("Hello, World!"))),
            Child(P(Text("Built with Blox"), Class("intro"))),
        )),
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
    Child(Div(
        Child(FormLabel(For("email"), Text("Email"))),
        Child(Input(
            InputType("email"),
            InputName("email"),
            Id("email"),
            Required(),
        )),
    )),
    Child(Div(
        Child(FormLabel(For("password"), Text("Password"))),
        Child(Input(
            InputType("password"),
            InputName("password"),
            Id("password"),
            Required(),
        )),
    )),
    Child(Button(
        ButtonType("submit"),
        Text("Login"),
    )),
)
```

### HTTP handlers

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    page := Html(
        Lang("en"),
        Child(Head(Child(HeadTitle(Text("Home"))))),
        Child(Body(
            Child(H1(Text("Welcome"))),
            Child(P(Text("This page was built with Blox"))),
        )),
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
├── options_content.go    # Text() and Child() helpers
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
