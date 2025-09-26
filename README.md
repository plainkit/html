# Plain: Type-Safe HTML in Go

**Why choose Plain?** Because building HTML should be as reliable as the rest of your Go code—with compile-time safety, IDE autocomplete, and zero runtime surprises.

## Why Plain?

### The Problem with Current Solutions

**Template engines fail at compile time**: Most Go HTML solutions parse templates at runtime, turning typos and structure errors into production bugs instead of build failures.

**Existing builders are verbose and fragile**: Struct-based HTML builders require tedious field assignments and don't prevent invalid attribute combinations like `<input href="/invalid">`.

**No real type safety**: You can pass any attribute to any element, discovering mistakes only when testing in a browser.

### The Plain Solution

Plain generates HTML using pure Go functions with compile-time validation. Each HTML element is a typed function that only accepts valid attributes and provides full IDE support.

```go
// This compiles and works
input := Input(
    AType("email"),
    AName("email"),
    ARequired(),
    APlaceholder("Enter email"),
)

// This fails at compile time - Href not valid for Input
input := Input(AHref("/invalid"))  // ❌ Compile error
```

## What is Plain?

Plain is a **function-based HTML component library** that generates HTML at compile time with zero runtime overhead.

- ✅ **Type-safe**: Each element only accepts valid attributes
- ✅ **Compile-time validation**: Catch HTML errors before deployment
- ✅ **Zero runtime cost**: Pure string building with no reflection
- ✅ **Full IDE support**: Autocomplete, go-to-definition, refactoring
- ✅ **Server-first**: Build complete HTML pages, not client-side components

### Core Philosophy

1. **HTML is data structures**: Represent your UI as composable Go functions
2. **Fail fast**: Invalid HTML should fail at compile time, not runtime
3. **Developer experience first**: Autocomplete, type checking, and refactoring should just work
4. **Zero magic**: Simple functions that build strings—nothing hidden

## How to Use Plain

### Installation

```bash
go get github.com/plainkit/html
```

### Quick Start

```go
package main

import (
    "fmt"
    . "github.com/plainkit/html"
)

func main() {
    page := Html(ALang("en"),
        Head(
            Title(T("My App")),
            Meta(ACharset("UTF-8")),
            Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
        ),
        Body(
            Header(
                H1(T("Welcome to Plain"), AClass("title")),
                Nav(
                    A(AHref("/about"), T("About")),
                    A(AHref("/contact"), T("Contact")),
                ),
            ),
            Main(
                P(T("Build type-safe HTML with Go functions."), AClass("intro")),
                Button(
                    AType("button"),
                    T("Get Started"),
                    AClass("btn btn-primary"),
                ),
            ),
        ),
    )

    fmt.Println("<!DOCTYPE html>")
    fmt.Println(Render(page))
}
```

### Building Forms

Forms are fully type-safe with element-specific attributes:

```go
func LoginForm() Component {
    return Form(
        AAction("/auth/login"),
        AMethod("POST"),
        AClass("login-form"),

        Div(AClass("field"),
            Label(AFor("email"), T("Email Address")),
            Input(
                AType("email"),
                AName("email"),
                AId("email"),
                ARequired(),
                APlaceholder("you@example.com"),
            ),
        ),

        Div(AClass("field"),
            Label(AFor("password"), T("Password")),
            Input(
                AType("password"),
                AName("password"),
                AId("password"),
                ARequired(),
                AMinlength("8"),
            ),
        ),

        Button(
            AType("submit"),
            T("Sign In"),
            AClass("btn-primary"),
        ),
    )
}
```

### HTTP Handlers

Plain components integrate seamlessly with Go's HTTP server:

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    page := Html(ALang("en"),
        Head(
            Title(T("Home")),
            Meta(ACharset("UTF-8")),
        ),
        Body(
            H1(T("Hello, World!")),
            P(T("Built with Plain")),
        ),
    )

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, "<!DOCTYPE html>\n")
    fmt.Fprint(w, Render(page))
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Component Composition

Build reusable components by composing smaller ones:

```go
func Card(title, content string, actions ...Component) Component {
    return Div(AClass("card"),
        Header(AClass("card-header"),
            H3(T(title), AClass("card-title")),
        ),
        Div(AClass("card-body"),
            P(T(content)),
        ),
        Footer(AClass("card-footer"),
            actions...,
        ),
    )
}

func Dashboard() Component {
    return Div(AClass("dashboard"),
        Card("Welcome", "Get started with Plain",
            Button(AType("button"), T("Learn More"), AClass("btn-outline")),
        ),
        Card("Stats", "View your analytics",
            A(AHref("/analytics"), T("View Details"), AClass("btn-primary")),
        ),
    )
}
```

### Working with CSS and JavaScript

Embed styles and scripts directly in your components:

```go
func StyledPage() Component {
    return Html(ALang("en"),
        Head(
            Title(T("Styled Page")),
            Style(T(`
                .hero {
                    background: linear-gradient(45deg, #667eea 0%, #764ba2 100%);
                    color: white;
                    padding: 2rem;
                    text-align: center;
                }
                .btn {
                    padding: 0.5rem 1rem;
                    border: none;
                    border-radius: 4px;
                    background: #4f46e5;
                    color: white;
                    cursor: pointer;
                }
            `)),
        ),
        Body(
            Div(AClass("hero"),
                H1(T("Beautiful Styling")),
                Button(
                    AType("button"),
                    T("Click Me"),
                    AClass("btn"),
                    AOnclick("alert('Hello from Plain!')"),
                ),
            ),
            Script(T(`
                console.log('Plain page loaded');
            `)),
        ),
    )
}
```

## Key Features

### 🔒 Type Safety

Each HTML element has its own argument interface that only accepts valid attributes:

```go
// ✅ Valid - these attributes work with Input
Input(AType("text"), AName("username"), ARequired())

// ✅ Valid - these attributes work with A
A(AHref("/home"), ATarget("_blank"), T("Home"))

// ❌ Invalid - compile error, Href not valid for Input
Input(AHref("/invalid"))

// ❌ Invalid - compile error, Required not valid for A
A(ARequired(), T("Link"))
```

### 🎯 IDE Support

- **Autocomplete**: Type `Input(` and see only valid options
- **Go to definition**: Jump to element implementations
- **Type checking**: Invalid combinations fail immediately
- **Refactoring**: Rename functions across your entire codebase

### 🚀 Zero Runtime Overhead

HTML generation uses method dispatch resolved at compile time:

```go
component := Div(AClass("test"), T("Hello"))
html := Render(component) // Pure string building, no reflection
```

### 🧪 Easy Testing

Components are just Go values—test them like any other Go code:

```go
func TestUserCard(t *testing.T) {
    card := UserCard("John Doe", "john@example.com")
    html := Render(card)

    assert.Contains(t, html, "John Doe")
    assert.Contains(t, html, "john@example.com")
    assert.Contains(t, html, `class="user-card"`)
}
```

### 📦 Complete HTML5 Support

Plain provides functions for all standard HTML5 elements:

- **Document**: `Html`, `Head`, `Body`, `Title`, `Meta`, `Link`
- **Sections**: `Header`, `Nav`, `Main`, `Section`, `Article`, `Aside`, `Footer`
- **Text**: `H1`-`H6`, `P`, `Div`, `Span`, `Strong`, `Em`, `Code`, `Pre`
- **Lists**: `Ul`, `Ol`, `Li`, `Dl`, `Dt`, `Dd`
- **Forms**: `Form`, `Input`, `Textarea`, `Button`, `Select`, `Option`, `Label`
- **Tables**: `Table`, `Thead`, `Tbody`, `Tr`, `Th`, `Td`
- **Media**: `Img`, `Video`, `Audio`, `Canvas`, `Svg`
- **And more**: All HTML5 elements with their specific attributes

### 🎨 SVG Support

Full SVG support with type-safe attributes:

```go
func Logo() Component {
    return Svg(
        Width("100"),
        Height("100"),
        ViewBox("0 0 100 100"),
        Circle(
            Cx("50"),
            Cy("50"),
            R("40"),
            Fill("#4f46e5"),
        ),
        Text(
            X("50"),
            Y("50"),
            TextAnchor("middle"),
            Fill("white"),
            T("LOGO"),
        ),
    )
}
```

## Architecture & Implementation

### Core Types

```go
// Component - interface implemented by all HTML elements
type Component interface {
    render(*strings.Builder)
}

// Node - represents an HTML element
type Node struct {
    Tag   string
    Attrs attrWriter
    Kids  []Component
    Void  bool // for self-closing tags
}

// TextNode - represents escaped text content
type TextNode string
```

### Global Attributes

Every HTML element can accept global attributes through the `Global` type:

```go
anyElement := Div(
    AId("main-content"),
    AClass("container active"),
    AData("role", "main"),
    AAria("label", "Main content area"),
    AStyle("margin-top: 1rem"),
    AOnclick("handleClick()"),
    // ... element-specific attributes
)
```

### File Organization

The codebase is organized for clarity and maintainability:

```
├── core_node.go         # Component interface, Node struct, Render()
├── core_global.go       # Global attributes (id, class, data-*, aria-*)
├── core_content.go      # Text() and Child() helpers
├── attrs.go             # All attribute option constructors
├── tag_div.go           # Div element and its specific attributes
├── tag_input.go         # Input element and its specific attributes
├── tag_form.go          # Form, Button, Select, etc.
├── tag_*.go             # One file per HTML element group
└── svg/                 # SVG elements in separate package
    ├── tag_circle.go
    ├── tag_rect.go
    └── ...
```

## Migration & Integration

### From Template Engines

**Before (html/template):**

```go
tmpl := `<div class="{{.Class}}">{{.Content}}</div>`
t := template.Must(template.New("div").Parse(tmpl))
t.Execute(w, map[string]interface{}{
    "Class": "container",
    "Content": "Hello World",
})
```

**After (Plain):**

```go
component := Div(AClass("container"), T("Hello World"))
fmt.Fprint(w, Render(component))
```

### From Other HTML Builders

**Before (gohtml or similar):**

```go
div := &gohtml.Div{
    Class: "container",
    Children: []gohtml.Element{
        &gohtml.Text{Content: "Hello World"},
    },
}
```

**After (Plain):**

```go
div := Div(AClass("container"), T("Hello World"))
```

### With Popular Frameworks

**With Gin:**

```go
r.GET("/", func(c *gin.Context) {
    page := Html(ALang("en"),
        Body(H1(T("Hello Gin + Plain!"))),
    )
    c.Header("Content-Type", "text/html")
    c.String(200, "<!DOCTYPE html>"+Render(page))
})
```

**With Echo:**

```go
e.GET("/", func(c echo.Context) error {
    page := Html(ALang("en"),
        Body(H1(T("Hello Echo + Plain!"))),
    )
    return c.HTML(200, "<!DOCTYPE html>"+Render(page))
})
```

**With htmx:**

```go
func TodoList(todos []Todo) Component {
    return Div(
        AId("todo-list"),
        ACustom("hx-get", "/todos"),
        ACustom("hx-trigger", "load"),
        ACustom("hx-swap", "innerHTML"),
        // Render todos...
    )
}
```

## Contributing

We welcome contributions! The codebase is designed to be approachable:

1. **Add new HTML elements**: Follow the pattern in existing `tag_*.go` files
2. **Improve type safety**: Add element-specific attributes and validation
3. **Enhance developer experience**: Better error messages, documentation, examples
4. **Test thoroughly**: Add tests for new elements and edge cases

Each element follows a consistent pattern:

- `*Attrs` struct for element-specific attributes
- `*Arg` interface for type constraints
- Constructor function accepting variadic arguments
- Apply methods connecting options to attributes

## License

MIT License - see LICENSE file for details.

---

**Ready to build type-safe HTML?** `go get github.com/plainkit/html` and start building reliable web applications with the confidence of Go's type system.
