# Examples

## Naming convention

Examples use a **range-based numbering** prefix by category:

| Range | Category                | Description                      |
| ----- | ----------------------- | -------------------------------- |
| `0xx` | **Core elements**       | Basic herald rendering features  |
| `1xx` | **Customization**       | Themes, palettes, options        |
| `2xx` | **Ecosystem / Pairing** | Integration with other libraries |

When adding a new example, pick the next available number in the matching range.

## Running

Some examples live in the root module and can be run directly with `go run`. Others are separate modules (they have their own `go.mod`) and need `cd` first.

### Core elements (`0xx`)

| Example                                     | Description                                                                       | Run                                      |
| ------------------------------------------- | --------------------------------------------------------------------------------- | ---------------------------------------- |
| [000_default-theme](000_default-theme/)     | All elements with the default Rose Pine theme                                     | `go run ./examples/000_default-theme/`   |
| [001_lists](001_lists/)                     | Flat, nested, mixed, and hierarchical lists                                       | `go run ./examples/001_lists/`           |
| [002_alerts](002_alerts/)                   | GitHub-style alert callouts (Note, Tip, Important, Warning, Caution)              | `go run ./examples/002_alerts/`          |
| [003_table](003_table/)                     | Table rendering: bordered, minimal, alignment, striped rows, captions, and footer | `go run ./examples/003_table/`           |
| [004_semantic-badges](004_semantic-badges/) | Semantic badges and tags with default and custom palettes                         | `go run ./examples/004_semantic-badges/` |
| [005_compose](005_compose/)                 | Compose multiple rendered blocks into a single output                             | `go run ./examples/005_compose/`         |

### Customization (`1xx`)

| Example                                       | Description                                                              | Run                                            |
| --------------------------------------------- | ------------------------------------------------------------------------ | ---------------------------------------------- |
| [100_custom-options](100_custom-options/)     | Override styles, decoration chars, and tokens via functional options     | `go run ./examples/100_custom-options/`        |
| [101_custom-palette](101_custom-palette/)     | Custom adaptive theme from 9 colors using `ColorPalette` and `LightDark` | `go run ./examples/101_custom-palette/`        |
| [102_builtin-themes](102_builtin-themes/)     | Built-in themes (Dracula, Catppuccin, Base16, Charm) matching huh        | `go run ./examples/102_builtin-themes/`        |
| [103_catppuccin-theme](103_catppuccin-theme/) | Build a full theme from the Catppuccin palette (separate module)         | `cd examples/103_catppuccin-theme && go run .` |

### Ecosystem / Pairing (`2xx`)

| Example                                                                       | Description                                                                 | Run                                                            |
| ----------------------------------------------------------------------------- | --------------------------------------------------------------------------- | -------------------------------------------------------------- |
| [200_chroma-syntax-highlighting](200_chroma-syntax-highlighting/)             | Plug in chroma for syntax-highlighted code blocks (separate module)         | `cd examples/200_chroma-syntax-highlighting && go run .`       |
| [201_tree-sitter-syntax-highlighting](201_tree-sitter-syntax-highlighting/)   | Plug in tree-sitter for AST-based syntax highlighting (separate module)     | `cd examples/201_tree-sitter-syntax-highlighting && go run .`  |
| [202_gotreesitter-syntax-highlighting](202_gotreesitter-syntax-highlighting/) | Pure-Go tree-sitter highlighting via gotreesitter (separate module)         | `cd examples/202_gotreesitter-syntax-highlighting && go run .` |
| [203_huh-form](203_huh-form/)                                                 | Using herald with huh for interactive TUI forms                             | `cd examples/203_huh-form && go run .`                         |
| [204_huh-wizard](204_huh-wizard/)                                             | Multi-step project scaffolder with herald + huh (separate module)           | `cd examples/204_huh-wizard && go run .`                       |
| [205_bubbletea-release-viewer](205_bubbletea-release-viewer/)                 | Scrollable release notes viewer with bubbletea viewport (separate module)   | `cd examples/205_bubbletea-release-viewer && go run .`         |
| [206_bubbletea-explorer](206_bubbletea-explorer/)                             | Sidebar + scrollable content pane explorer with bubbletea (separate module) | `cd examples/206_bubbletea-explorer && go run .`               |
| [207_tview-explorer](207_tview-explorer/)                                     | Sidebar + scrollable content pane explorer with tview (separate module)     | `cd examples/207_tview-explorer && go run .`                   |
