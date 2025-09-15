# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

description: Guidelines for creating Blox idiomatic UI components
globs: "ui/\*.go"
alwaysApply: true

---

## Project Overview

Blox is a function-based HTML component library for Go that generates HTML at compile time with zero runtime overhead. It provides type-safe HTML construction through function composition instead of template engines or struct-based builders.

## Development Commands

### Core Library Development

- **Build**: `go build` (builds the main library)
- **Test**: `go test ./...` (run all tests, though no test files exist currently)
- **Run demo**: `cd demo && go run main.go` (starts demo server on :3000)

### Demo Application

- **Start demo server**: `cd demo && go run main.go`
- **Build CSS**: CSS is embedded in `css/embed.go` - Tailwind CSS compilation happens externally

### Module Structure

- Main library: `github.com/plainkit/blox` (Go 1.20+)
- Demo app: `github.com/plainkit/blox/demo` (Go 1.21+, replaces main library locally)

## Core Architecture

### Component System

- **Component**: Interface with `render(*strings.Builder)` method
- **Node**: Core struct representing HTML elements (tag + attributes + children)
- **TextNode**: Represents escaped text content
- **Global**: Universal options that work with any HTML element

### File Organization Pattern

```
core_*.go           # Core interfaces and rendering system
options_*.go        # Option constructors (Text, Child, Class, etc.)
tag_*.go           # HTML element implementations (one logical group per file)
ui/*.go            # Higher-level UI components with styling
css/               # CSS embedding and Tailwind integration
demo/              # Example application showcasing the library
```

### HTML Element Architecture

Each `tag_*.go` file follows this pattern:

1. **Attrs struct**: Element-specific attributes (e.g., `DivAttrs`, `InputAttrs`)
2. **Arg interface**: Type constraint for element options (e.g., `DivArg`, `InputArg`)
3. **Constructor function**: Main element function accepting variadic args
4. **Option types**: Specific attribute setters (e.g., `InputType()`, `Required()`)
5. **Apply methods**: Connect options to attributes via interface satisfaction
6. **writeAttrs method**: Generates HTML attribute strings

### UI Components System

The `ui/` folder contains higher-level components following this pattern:

#### Interface Adapter Pattern

```go
// Accept both UI-specific and core Blox arguments
func Component(args ...interface{}) x.Component {
    state := &componentState{}

    for _, arg := range args {
        if adapted := adaptComponentArg(arg); adapted != nil {
            adapted.applyUIComponent(state)
        }
    }

    // Apply styling and build final component
}
```

#### Key UI Principles

- **No Custom Child Functions**: Always use `x.Child()`, never `ComponentChild()`
- **No Custom Text Functions**: Always use `x.Text()`, never `ComponentText()`
- **Interface Flexibility**: Support both `Component(x.Id("test"))` and UI-specific options
- **Semantic CSS**: Use design system tokens (`bg-card`, `text-muted-foreground`, etc.)

## Type Safety System

### Compile-Time Guarantees

- Each HTML element only accepts valid attributes for that element
- Invalid combinations like `Input(Href("/invalid"))` fail at compile time
- Children and text content are type-safe through the Component interface

### Attribute Validation

- Global attributes (class, id, etc.) work with any element via `Global` type
- Element-specific attributes are constrained by their respective `Arg` interfaces
- Method dispatch resolves at compile time with zero runtime overhead

## CSS Integration

### Tailwind CSS System

- CSS is compiled externally and embedded in `css/embed.go`
- Demo serves CSS via `/assets/styles.css` endpoint
- Uses CSS custom properties with `@theme inline` for design tokens
- System dark mode detection via `@media (prefers-color-scheme: dark)`

### Design System

- Semantic color tokens: `bg-card`, `text-card-foreground`, `text-muted-foreground`
- Consistent spacing and sizing through Tailwind utilities
- Component variants follow modern UI library patterns

## Adding New Elements

### Core HTML Elements

1. Create `tag_newelement.go` following the established pattern
2. Define element-specific `Attrs` struct and `Arg` interface
3. Implement constructor function with variadic arguments
4. Add element-specific option constructors and apply methods
5. Implement `writeAttrs()` method for HTML generation

### UI Components

1. Create file in `ui/` directory following interface adapter pattern
2. Use semantic CSS classes for styling
3. Ensure compatibility with core Blox arguments via adapter system
4. Follow established naming conventions (`CardHeader`, `CardContent`, etc.)

## Demo Application Structure

The demo showcases practical usage patterns:

- **Layout system**: Consistent navigation and footer across pages
- **Component composition**: Cards, buttons, forms using the UI library
- **HTTP integration**: Standard library handlers serving Blox-generated HTML
- **CSS serving**: Embedded Tailwind CSS served as static asset

Key demo patterns to follow when adding examples:

- Use semantic HTML structure
- Demonstrate type safety benefits
- Show composition patterns
- Include responsive design principles

## UI Component Implementation Template

When creating UI components in the `ui/` folder, follow this pattern:

```go
package ui

import x "github.com/plainkit/blox"

// ComponentArg interface for UI component arguments
type ComponentArg interface {
    applyUIComponent(*componentState)
}

// componentState holds configuration
type componentState struct {
    children []x.Component
    baseArgs []x.DivArg // or appropriate tag args
}

// Universal adapter for core blox args
type ComponentArgAdapter struct{ arg x.DivArg }

func (a ComponentArgAdapter) applyUIComponent(s *componentState) {
    s.baseArgs = append(s.baseArgs, a.arg)
}

// Interface adapter magic
func adaptComponentArg(arg interface{}) ComponentArg {
    if uiArg, ok := arg.(ComponentArg); ok {
        return uiArg
    }
    if coreArg, ok := arg.(x.DivArg); ok {
        return ComponentArgAdapter{coreArg}
    }
    return nil
}

// Main component function
func Component(args ...interface{}) x.Component {
    state := &componentState{}

    for _, arg := range args {
        if adapted := adaptComponentArg(arg); adapted != nil {
            adapted.applyUIComponent(state)
        }
    }

    // Apply base styling classes
    classes := "component-base-classes semantic-tokens"
    componentArgs := append([]x.DivArg{x.Class(classes)}, state.baseArgs...)

    // Add children
    for _, child := range state.children {
        componentArgs = append(componentArgs, x.Child(child))
    }

    return x.Div(componentArgs...)
}
```

This pattern ensures excellent DX while maintaining Blox architectural consistency.
