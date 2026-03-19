<h1 align="center">
  herald
</h1>

<h2 align="center" style="font-size: 1.5rem;">
    HTML-inspired typography for terminal UIs in Go.
</h2>

<p align="center">
  <a href="https://github.com/indaco/herald/actions/workflows/ci.yml" target="_blank">
    <img src="https://github.com/indaco/herald/actions/workflows/ci.yml/badge.svg" alt="CI" />
  </a>
  <a href="https://codecov.io/gh/indaco/herald" target="_blank">
    <img src="https://codecov.io/gh/indaco/herald/branch/main/graph/badge.svg" alt="Code coverage" />
  </a>
  <a href="https://goreportcard.com/report/github.com/indaco/herald" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/herald" alt="Go Report Card" />
  </a>
  <a href="https://github.com/indaco/herald/security" target="_blank">
    <img src="https://img.shields.io/badge/security-govulncheck-green" alt="Security Scan" />
  </a>
  <a href="https://github.com/indaco/herald/releases" target="_blank">
    <img src="https://img.shields.io/github/v/tag/indaco/herald?label=version&sort=semver&color=4c1" alt="version">
  </a>
  <a href="https://pkg.go.dev/github.com/indaco/herald" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/indaco/herald.svg" alt="Go Reference" />
  </a>
  <a href="LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square" alt="License" />
  </a>
  <a href="https://www.jetify.com/devbox" target="_blank">
    <img src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

<p align="center">
  <b><a href="#quick-start">Quick Start</a></b> |
  <b><a href="#available-elements">Elements</a></b> |
  <b><a href="#customization">Customization</a></b> |
  <b><a href="#built-in-themes">Themes</a></b> |
  <b><a href="#examples">Examples</a></b>
</p>

Herald maps familiar HTML elements (H1–H6, P, Blockquote, UL, OL, Code, HR, Alerts, and inline styles) to styled terminal output, built on [lipgloss v2](https://github.com/charmbracelet/lipgloss).

It ships with a Rose Pine-inspired default theme, built-in themes matching the Charm ecosystem (Dracula, Catppuccin, Base16, Charm) for seamless pairing with [huh](https://github.com/charmbracelet/huh) and other Charm-based TUIs, and full style customization via functional options and `ColorPalette`.

<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-hero.png" alt="herald demo output" width="600" />
</p>

<p align="center"><em>Default Rose Pine theme (dark and light). Herald also ships with Dracula, Catppuccin, Base16, and Charm themes.</em></p>

## Installation

```sh
go get github.com/indaco/herald
```

Requires Go 1.25 or later.

## Quick start

```go
package main

import (
    "fmt"
    "github.com/indaco/herald"
)

func main() {
    ty := herald.New()

    fmt.Println(ty.H1("Getting Started"))
    fmt.Println(ty.P("Herald renders terminal typography using lipgloss styles."))
    fmt.Println(ty.UL("Headings", "Block elements", "Inline styles"))
}
```

## Available elements

### Headings

<details>
<summary><b>Preview</b></summary>
<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-headings.png" alt="headings demo" width="600" />
</p>
</details>

H1–H3 render with a repeated underline character beneath the text. H4–H6 render with a left bar prefix.

| Method     | Decoration | Default character |
| ---------- | ---------- | ----------------- |
| `H1(text)` | underline  | `═`               |
| `H2(text)` | underline  | `─`               |
| `H3(text)` | underline  | `·`               |
| `H4(text)` | bar prefix | `▎`               |
| `H5(text)` | bar prefix | `▎`               |
| `H6(text)` | bar prefix | `▎`               |

### Block elements

<details>
<summary><b>Preview</b></summary>
<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-blocks.png" alt="blocks demo" width="600" />
</p>
</details>

| Method                  | Description                                                                                 |
| ----------------------- | ------------------------------------------------------------------------------------------- |
| `P(text)`               | Paragraph                                                                                   |
| `Blockquote(text)`      | Indented block with a left bar; supports multi-line input                                   |
| `Code(text, lang)`      | Inline code with background highlight; `lang` is optional, used when a CodeFormatter is set |
| `CodeBlock(text, lang)` | Fenced code block with padding; optional line numbers and syntax highlighting               |
| `HR()`                  | Horizontal rule, configurable width and character                                           |
| `DL(pairs)`             | Definition list from `[][2]string` pairs (term, description)                                |
| `DT(text)`              | Definition term (standalone)                                                                |
| `DD(text)`              | Definition description (standalone)                                                         |

```go
fmt.Println(ty.Blockquote("First line.\nSecond line."))
fmt.Println(ty.Code("os.Exit(1)"))
fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"hello\")\n}"))
fmt.Println(ty.HR())

fmt.Println(ty.DL([][2]string{
    {"Go", "A statically typed, compiled language"},
    {"Rust", "A systems programming language"},
}))
```

### Lists

<details>
<summary><b>Preview</b></summary>
<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-lists.png" alt="lists demo" width="600" />
</p>
</details>

```go
fmt.Println(ty.UL("Apples", "Bananas", "Cherries"))
fmt.Println(ty.OL("First", "Second", "Third"))
```

`UL` uses a bullet character (default `•`). `OL` uses `1.`, `2.`, `3.` markers.

### Nested lists

`NestUL` and `NestOL` render hierarchical lists with configurable indentation, bullet cycling, and mixed nesting.

**Constructors:**

| Function                                | Description                         |
| --------------------------------------- | ----------------------------------- |
| `Item(text)`                            | Leaf item (no children)             |
| `Items(texts...)`                       | Batch-convert strings to leaf items |
| `ItemWithChildren(text, children...)`   | Item with unordered sub-list        |
| `ItemWithOLChildren(text, children...)` | Item with ordered sub-list          |

```go
// Nested unordered list with mixed sub-lists
fmt.Println(ty.NestUL(
    herald.Item("Fruits"),
    herald.ItemWithChildren("Vegetables",
        herald.Item("Carrots"),
        herald.Item("Peas"),
    ),
    herald.ItemWithOLChildren("Ranked Desserts",
        herald.Item("Ice cream"),
        herald.Item("Cake"),
    ),
))
```

```text
• Fruits
• Vegetables
  ◦ Carrots
  ◦ Peas
• Ranked Desserts
  1. Ice cream
  2. Cake
```

```go
// Nested ordered list
fmt.Println(ty.NestOL(
    herald.Item("Introduction"),
    herald.ItemWithOLChildren("Main Topics",
        herald.Item("Architecture"),
        herald.Item("Design"),
    ),
    herald.Item("Conclusion"),
))
```

```text
1. Introduction
2. Main Topics
  1. Architecture
  2. Design
3. Conclusion
```

Enable `WithHierarchicalNumbers(true)` for outline-style numbering (`2.1.`, `2.2.`):

```go
ty := herald.New(herald.WithHierarchicalNumbers(true))

fmt.Println(ty.NestOL(
    herald.Item("Introduction"),
    herald.ItemWithOLChildren("Main Topics",
        herald.Item("Architecture"),
        herald.Item("Design"),
    ),
    herald.Item("Conclusion"),
))
```

```text
1. Introduction
2. Main Topics
  2.1. Architecture
  2.2. Design
3. Conclusion
```

### Inline styles

<details>
<summary><b>Preview</b></summary>
<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-inline.png" alt="inline styles demo" width="600" />
</p>
</details>

| Method                | Renders as                                                                  |
| --------------------- | --------------------------------------------------------------------------- |
| `Bold(text)`          | Bold                                                                        |
| `Italic(text)`        | Italic                                                                      |
| `Underline(text)`     | Underlined                                                                  |
| `Strikethrough(text)` | Strikethrough                                                               |
| `Small(text)`         | Faint                                                                       |
| `Mark(text)`          | Highlighted background                                                      |
| `Link(label, url)`    | Styled link; `url` is optional — when both differ, renders as `label (url)` |
| `Kbd(text)`           | Keyboard key indicator                                                      |
| `Abbr(abbr, desc)`    | Abbreviation; `desc` is optional, appended in parentheses                   |
| `Sub(text)`           | Subscript, prefixed with `_`                                                |
| `Sup(text)`           | Superscript, prefixed with `^`                                              |

```go
fmt.Println(ty.Bold("important") + " and " + ty.Italic("nuanced"))
fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
fmt.Println(ty.Link("Go website", "https://go.dev"))
fmt.Println(ty.Abbr("CSS", "Cascading Style Sheets"))
fmt.Println(ty.Sub("2") + "O" + ty.Sup("n"))
```

### Alerts

<details>
<summary><b>Preview</b></summary>
<p align="center">
  <img src="https://raw.githubusercontent.com/indaco/gh-assets/main/herald/demo-alerts.png" alt="alerts demo" width="600" />
</p>
</details>

GitHub-style alert callouts with colored bars, icons, and labels. Five types are supported: Note, Tip, Important, Warning, and Caution.

| Method            | Icon | Color  | Description                       |
| ----------------- | ---- | ------ | --------------------------------- |
| `Note(text)`      | `○`  | Blue   | Useful information for the reader |
| `Tip(text)`       | `▸`  | Green  | Helpful advice                    |
| `Important(text)` | `‼`  | Purple | Key information                   |
| `Warning(text)`   | `⚠`  | Amber  | Urgent attention needed           |
| `Caution(text)`   | `◇`  | Red    | Risk or negative outcomes         |

```go
fmt.Println(ty.Note("Useful information that users should know."))
fmt.Println(ty.Tip("Helpful advice for doing things better."))
fmt.Println(ty.Important("Key information users need to know."))
fmt.Println(ty.Warning("Urgent info that needs immediate attention."))
fmt.Println(ty.Caution("Advises about risks or negative outcomes."))
```

```text
│ ○ Note
│ Useful information that users should know.

│ ⚠ Warning
│ Urgent info that needs immediate attention.
```

You can also use the generic `Alert` method with an `AlertType`:

```go
fmt.Println(ty.Alert(herald.AlertNote, "Generic alert call."))
```

## Customization

### Functional options

Pass options to `herald.New()` to override individual styles or tokens.

```go
ty := herald.New(
    herald.WithHRWidth(60),
    herald.WithBulletChar("-"),
    herald.WithH1Style(
        lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000")),
    ),
)
```

**Style options** — each accepts a `lipgloss.Style`:

| Option                        | Targets                 |
| ----------------------------- | ----------------------- |
| `WithH1Style` – `WithH6Style` | Heading levels 1–6      |
| `WithParagraphStyle`          | `P`                     |
| `WithBlockquoteStyle`         | `Blockquote`            |
| `WithCodeInlineStyle`         | `Code`                  |
| `WithCodeBlockStyle`          | `CodeBlock`             |
| `WithCodeLineNumberStyle`     | Code block line numbers |
| `WithHRStyle`                 | `HR`                    |
| `WithBoldStyle`               | `Bold`                  |
| `WithItalicStyle`             | `Italic`                |
| `WithUnderlineStyle`          | `Underline`             |
| `WithStrikethroughStyle`      | `Strikethrough`         |
| `WithSmallStyle`              | `Small`                 |
| `WithMarkStyle`               | `Mark`                  |
| `WithLinkStyle`               | `Link`                  |
| `WithKbdStyle`                | `Kbd`                   |
| `WithAbbrStyle`               | `Abbr`                  |
| `WithListBulletStyle`         | Bullet/number marker    |
| `WithListItemStyle`           | List item text          |
| `WithDTStyle`                 | Definition term         |
| `WithDDStyle`                 | Definition description  |
| `WithAlertStyle(type, style)` | Alert of given type     |

**Token options** — each accepts a `string` or `int`:

| Option                         | Default            | Description                                           |
| ------------------------------ | ------------------ | ----------------------------------------------------- |
| `WithH1UnderlineChar(c)`       | `═`                | Underline character for H1                            |
| `WithH2UnderlineChar(c)`       | `─`                | Underline character for H2                            |
| `WithH3UnderlineChar(c)`       | `·`                | Underline character for H3                            |
| `WithHeadingBarChar(c)`        | `▎`                | Bar prefix character for H4–H6                        |
| `WithBulletChar(c)`            | `•`                | Bullet character for `UL`                             |
| `WithNestedBulletChars(chars)` | `•`, `◦`, `▪`, `▹` | Bullet characters cycling per depth for `NestUL`      |
| `WithListIndent(n)`            | `2`                | Spaces per nesting level for `NestUL`/`NestOL`        |
| `WithHierarchicalNumbers(b)`   | `false`            | Outline-style numbering for nested `OL` (e.g. `2.1.`) |
| `WithHRChar(c)`                | `─`                | Character repeated for `HR`                           |
| `WithHRWidth(w)`               | `40`               | Width of `HR` in characters                           |
| `WithBlockquoteBar(c)`         | `│`                | Left bar character for `Blockquote`                   |
| `WithCodeLineNumbers(b)`       | `false`            | Show line numbers in `CodeBlock`                      |
| `WithCodeLineNumberSep(c)`     | `│`                | Separator between line numbers and code               |
| `WithAlertBar(c)`              | `│`                | Left bar character for alerts                         |
| `WithAlertIcon(type, icon)`    | per-type           | Override icon for a specific alert type               |
| `WithAlertLabel(type, label)`  | per-type           | Override label for a specific alert type              |

### Code formatting

`WithCodeFormatter` accepts a `func(code, language string) string` callback. When set, `Code()` and `CodeBlock()` pass the raw text and language string to the formatter before applying the lipgloss style. When not set (the default), behavior is unchanged.

```go
import (
    "strings"

    "github.com/alecthomas/chroma/v2/quick"
    "github.com/indaco/herald"
)

func chromaFormatter(style string) func(code, language string) string {
    return func(code, language string) string {
        var buf strings.Builder
        err := quick.Highlight(&buf, code, language, "terminal256", style)
        if err != nil {
            return code
        }
        return strings.TrimRight(buf.String(), "\n")
    }
}

ty := herald.New(
    herald.WithCodeFormatter(chromaFormatter("catppuccin-mocha")),
)

fmt.Println(ty.CodeBlock(`func main() { fmt.Println("hello") }`, "go"))
```

See [`examples/07_syntax-highlighting/`](examples/07_syntax-highlighting/) for a chroma-based example, or [`examples/08_tree-sitter-highlighting/`](examples/08_tree-sitter-highlighting/) for a tree-sitter-based alternative.

### Line numbers in code blocks

Enable line numbers with `WithCodeLineNumbers(true)`. Line numbers are right-aligned, styled with `CodeLineNumber` (defaults to the `Muted` palette color), and separated from code by `CodeLineNumberSep` (default `│`). Line numbers are added after the `CodeFormatter` runs, so they work with syntax highlighting.

```go
ty := herald.New(
    herald.WithCodeLineNumbers(true),
)

fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"hello\")\n}", "go"))
```

```text
1│ func main() {
2│     fmt.Println("hello")
3│ }
```

Customize the separator and style:

```go
ty := herald.New(
    herald.WithCodeLineNumbers(true),
    herald.WithCodeLineNumberSep(":"),
    herald.WithCodeLineNumberStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))),
)
```

### Color palette

`ColorPalette` lets you define 9 colors and derive a complete theme from them. All style fields map from this palette; configurable tokens use the same defaults as `DefaultTheme`.

| Field       | Maps to                                              |
| ----------- | ---------------------------------------------------- |
| `Primary`   | H1 headings                                          |
| `Secondary` | H2, list bullets                                     |
| `Tertiary`  | H3, links                                            |
| `Accent`    | H4, mark background                                  |
| `Highlight` | H5, `Abbr`                                           |
| `Muted`     | H6, blockquote, HR, sub/sup, `DD`                    |
| `Text`      | Body text, paragraphs, list items, inline code, `DT` |
| `Surface`   | Background for inline code and `Kbd`                 |
| `Base`      | Background for code blocks; mark foreground          |

Pass the palette to `New()` via `WithPalette`, or call `ThemeFromPalette` to construct a `Theme` value directly.

```go
// Dracula-inspired palette
palette := herald.ColorPalette{
    Primary:   lipgloss.Color("#bd93f9"), // purple
    Secondary: lipgloss.Color("#ff79c6"), // pink
    Tertiary:  lipgloss.Color("#8be9fd"), // cyan
    Accent:    lipgloss.Color("#ffb86c"), // orange
    Highlight: lipgloss.Color("#ff5555"), // red
    Muted:     lipgloss.Color("#6272a4"), // comment gray
    Text:      lipgloss.Color("#f8f8f2"), // foreground
    Surface:   lipgloss.Color("#44475a"), // current line
    Base:      lipgloss.Color("#282a36"), // background
}

ty := herald.New(herald.WithPalette(palette))
```

#### Alert palette

`AlertPalette` lets you override the 5 alert colors independently from the main `ColorPalette`:

```go
ty := herald.New(
    herald.WithAlertPalette(herald.AlertPalette{
        Note:      lipgloss.Color("#0969DA"),
        Tip:       lipgloss.Color("#1A7F37"),
        Important: lipgloss.Color("#8250DF"),
        Warning:   lipgloss.Color("#9A6700"),
        Caution:   lipgloss.Color("#CF222E"),
    }),
)
```

Individual alert icons and labels can also be customized:

```go
ty := herald.New(
    herald.WithAlertIcon(herald.AlertTip, "💡"),
    herald.WithAlertLabel(herald.AlertNote, "Info"),
)
```

You can combine `WithPalette` with other options to override specific fields after the palette is applied:

```go
ty := herald.New(
    herald.WithPalette(palette),
    herald.WithHRWidth(60),
    herald.WithBulletChar("-"),
)
```

### Built-in themes

Herald ships with named themes that match [huh](https://charm.land/huh)'s built-in color palettes. Colors auto-adapt to light/dark terminal backgrounds using `lipgloss.HasDarkBackground`.

| Function            | Palette                                                           |
| ------------------- | ----------------------------------------------------------------- |
| `DraculaTheme()`    | [Dracula](https://draculatheme.com)                               |
| `CatppuccinTheme()` | [Catppuccin](https://catppuccin.com) Mocha (dark) / Latte (light) |
| `Base16Theme()`     | ANSI base16 terminal colors                                       |
| `CharmTheme()`      | [Charm](https://charm.sh) brand colors                            |

```go
ty := herald.New(herald.WithTheme(herald.DraculaTheme()))
```

Pair with the corresponding huh theme for consistent styling across forms and typography:

```go
form := huh.NewForm(...).WithTheme(huh.ThemeDracula())
ty := herald.New(herald.WithTheme(herald.DraculaTheme()))
```

### Custom theme

The easiest way to customize is to start from an existing theme and modify specific fields:

```go
custom := herald.DefaultTheme()
custom.H1 = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000")).MarginBottom(1)
custom.BulletChar = "-"

ty := herald.New(herald.WithTheme(custom))
```

For a fully custom theme, construct a `Theme` struct directly:

```go
custom := herald.Theme{
    H1:        lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF")),
    H2:        lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#AAAAAA")),
    Paragraph: lipgloss.NewStyle().MarginBottom(1),
    // set remaining Theme fields as needed...

    H1UnderlineChar: "=",
    H2UnderlineChar: "-",
    H3UnderlineChar: ".",
    HeadingBarChar:  ">",
    BulletChar:      "*",
    HRChar:          "-",
    HRWidth:         40,
    BlockquoteBar:   "|",
}

ty := herald.New(herald.WithTheme(custom))
```

## Examples

Runnable examples are in the [`examples/`](examples/) directory:

| Example                                                              | Description                                                                                | Run                                                   |
| -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | ----------------------------------------------------- |
| [00_default-theme](examples/00_default-theme/)                       | All elements with the default Rose Pine theme                                              | `go run ./examples/00_default-theme/`                 |
| [01_lists](examples/01_lists/)                                       | Flat, nested, mixed, and hierarchical lists                                                | `go run ./examples/01_lists/`                         |
| [02_alerts](examples/02_alerts/)                                     | GitHub-style alert callouts (Note, Tip, Important, Warning, Caution)                       | `go run ./examples/02_alerts/`                        |
| [03_custom-options](examples/03_custom-options/)                     | Override styles, decoration chars, and tokens via functional options                       | `go run ./examples/03_custom-options/`                |
| [04_palette](examples/04_palette/)                                   | Generate a full theme from 9 colors using `ColorPalette`                                   | `go run ./examples/04_palette/`                       |
| [05_builtin-themes](examples/05_builtin-themes/)                     | Built-in themes (Dracula, Catppuccin, Base16, Charm) matching huh                          | `go run ./examples/05_builtin-themes/`                |
| [06_catppuccin-theme](examples/06_catppuccin-theme/)                 | Build a full theme from the [Catppuccin](https://catppuccin.com) palette (separate module) | `cd examples/06_catppuccin-theme && go run .`         |
| [07_syntax-highlighting](examples/07_syntax-highlighting/)           | Plug in chroma for syntax-highlighted code blocks (separate module)                        | `cd examples/07_syntax-highlighting && go run .`      |
| [08_tree-sitter-highlighting](examples/08_tree-sitter-highlighting/) | Plug in tree-sitter for AST-based syntax highlighting (separate module)                    | `cd examples/08_tree-sitter-highlighting && go run .` |

The _06_catppuccin-theme_, _07_syntax-highlighting_, and _08_tree-sitter-highlighting_ examples each have their own `go.mod` to keep external dependencies out of herald's core module.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
