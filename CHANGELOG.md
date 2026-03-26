# Changelog

All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). The changelog is generated and managed by [sley](https://github.com/indaco/sley).

## v0.8.0 - 2026-03-26

### 🚀 Enhancements

- add semantic palette with themed badge and tag variants ([d5d2045](https://github.com/indaco/herald/commit/d5d2045)) ([#44](https://github.com/indaco/herald/pull/44))

### 📖 Documentation

- **README:** use go get @latest in installation instructions ([b1ba5ef](https://github.com/indaco/herald/commit/b1ba5ef))
- **README:** fix ecosystem pairing sections and add tview ([7880674](https://github.com/indaco/herald/commit/7880674))
- **README:** add Pairing with bubbletea section ([27b5f65](https://github.com/indaco/herald/commit/27b5f65)) ([#40](https://github.com/indaco/herald/pull/40))

### 🏡 Chores

- refactor input and resize for model state in bubbletea examples ([b010abc](https://github.com/indaco/herald/commit/b010abc)) ([#43](https://github.com/indaco/herald/pull/43))
- **examples:** add bubbletea and tview explorer demos ([f42d44e](https://github.com/indaco/herald/commit/f42d44e)) ([#42](https://github.com/indaco/herald/pull/42))
- **examples:** reorder with range-based numbering convention ([5f9864a](https://github.com/indaco/herald/commit/5f9864a)) ([#41](https://github.com/indaco/herald/pull/41))

**Full Changelog:** [v0.7.0...v0.8.0](https://github.com/indaco/herald/compare/v0.7.0...v0.8.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.7.0 - 2026-03-23

### 🚀 Enhancements

- add KV and KVGroup for key-value pair rendering ([5d0d00c](https://github.com/indaco/herald/commit/5d0d00c)) ([#35](https://github.com/indaco/herald/pull/35))
- add HRWithLabel for labeled horizontal separators ([9263911](https://github.com/indaco/herald/commit/9263911)) ([#23](https://github.com/indaco/herald/pull/23))

### 🩹 Fixes

- use TableBorderSet.ColumnSep instead of hardcoded pipe ([674be75](https://github.com/indaco/herald/commit/674be75)) ([#33](https://github.com/indaco/herald/pull/33))
- add depth limit to nested list rendering ([01bb8e4](https://github.com/indaco/herald/commit/01bb8e4)) ([#32](https://github.com/indaco/herald/pull/32))
- add upper bound validation on HRWidth and FootnoteDividerWidth ([646277f](https://github.com/indaco/herald/commit/646277f)) ([#28](https://github.com/indaco/herald/pull/28))
- use display width for heading underlines and deduplicate style call ([130d192](https://github.com/indaco/herald/commit/130d192)) ([#25](https://github.com/indaco/herald/pull/25))
- guard against empty NestedBulletChars and protect default slice ([95e5327](https://github.com/indaco/herald/commit/95e5327)) ([#24](https://github.com/indaco/herald/pull/24))

### 💅 Refactors

- cache terminal background detection and reduce option boilerplate ([c95d8c6](https://github.com/indaco/herald/commit/c95d8c6)) ([#26](https://github.com/indaco/herald/pull/26))

### 📖 Documentation

- fix Color palette documentation and example for adaptive colors ([3b1fe85](https://github.com/indaco/herald/commit/3b1fe85)) ([#37](https://github.com/indaco/herald/pull/37))
- **README:** add global padding and framing composition recipe ([629df06](https://github.com/indaco/herald/commit/629df06)) ([#36](https://github.com/indaco/herald/pull/36))

### ✅ Tests

- add edge case coverage tests ([760335e](https://github.com/indaco/herald/commit/760335e)) ([#34](https://github.com/indaco/herald/pull/34))
- add direct tests for BoxBorderSet and DefaultAlertConfigs ([2d4769e](https://github.com/indaco/herald/commit/2d4769e)) ([#30](https://github.com/indaco/herald/pull/30))
- assert correct output in table concurrency test ([197ce54](https://github.com/indaco/herald/commit/197ce54)) ([#29](https://github.com/indaco/herald/pull/29))
- assert style options are actually applied, not just panic-free ([b903fc1](https://github.com/indaco/herald/commit/b903fc1)) ([#27](https://github.com/indaco/herald/pull/27))

### 🏡 Chores

- **examples:** add KVGroup to hero demo ([8368862](https://github.com/indaco/herald/commit/8368862)) ([#39](https://github.com/indaco/herald/pull/39))
- **examples:** add bubbletea integration demos ([a4fddc8](https://github.com/indaco/herald/commit/a4fddc8)) ([#38](https://github.com/indaco/herald/pull/38))

### 🤖 CI

- harden GitHub Actions workflows with zizmor recommendations ([cf8b8fb](https://github.com/indaco/herald/commit/cf8b8fb)) ([#31](https://github.com/indaco/herald/pull/31))

**Full Changelog:** [v0.6.0...v0.7.0](https://github.com/indaco/herald/compare/v0.6.0...v0.7.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.6.0 - 2026-03-22

### 🚀 Enhancements

- add Footnote typography element ([bb876b4](https://github.com/indaco/herald/commit/bb876b4)) ([#19](https://github.com/indaco/herald/pull/19))
- add Badge and Tag inline typography elements ([a5cb277](https://github.com/indaco/herald/commit/a5cb277)) ([#18](https://github.com/indaco/herald/pull/18))
- add Address and AddressCard block typography elements ([e0548e1](https://github.com/indaco/herald/commit/e0548e1)) ([#17](https://github.com/indaco/herald/pull/17))
- add Ins/Del inline typography elements ([0957853](https://github.com/indaco/herald/commit/0957853)) ([#16](https://github.com/indaco/herald/pull/16))

### 🩹 Fixes

- separate blockquote bar style for visible rendering ([ef53ddf](https://github.com/indaco/herald/commit/ef53ddf)) ([#20](https://github.com/indaco/herald/pull/20))

### 📖 Documentation

- **README:** move Color palette section under Themes ([285f3a0](https://github.com/indaco/herald/commit/285f3a0)) ([#22](https://github.com/indaco/herald/pull/22))
- **README:** restructure - add composition patterns, split options tables ([2eb9a3a](https://github.com/indaco/herald/commit/2eb9a3a)) ([#21](https://github.com/indaco/herald/pull/21))

### 🏡 Chores

- **examples:** add basic table example in hero demo ([d78f585](https://github.com/indaco/herald/commit/d78f585))

**Full Changelog:** [v0.5.0...v0.6.0](https://github.com/indaco/herald/compare/v0.5.0...v0.6.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.5.0 - 2026-03-20

### 🚀 Enhancements

- add table rendering ([4160ab7](https://github.com/indaco/herald/commit/4160ab7)) ([#12](https://github.com/indaco/herald/pull/12))

### 📖 Documentation

- **README:** reorganize sections and fix consistency issues ([63888f6](https://github.com/indaco/herald/commit/63888f6)) ([#13](https://github.com/indaco/herald/pull/13))
- **README:** updated with improved theme and customization details ([42a5228](https://github.com/indaco/herald/commit/42a5228))
- add Pairing with huh section and example ([ca290ab](https://github.com/indaco/herald/commit/ca290ab)) ([#11](https://github.com/indaco/herald/pull/11))

### 🏡 Chores

- **examples:** add example with github.com/odvcencio/gotreesitter ([6123760](https://github.com/indaco/herald/commit/6123760)) ([#15](https://github.com/indaco/herald/pull/15))
- **examples:** rename syntax highlighting examples for consistency ([21a0350](https://github.com/indaco/herald/commit/21a0350)) ([#14](https://github.com/indaco/herald/pull/14))
- add charm/huh example and update README ([82beef9](https://github.com/indaco/herald/commit/82beef9)) ([#10](https://github.com/indaco/herald/pull/10))
- update pre-push githook to use test-race recipe ([67fd08f](https://github.com/indaco/herald/commit/67fd08f))
- update devbox-init.sh available commands to match justfile recipes ([8569017](https://github.com/indaco/herald/commit/8569017))

**Full Changelog:** [v0.4.0...v0.5.0](https://github.com/indaco/herald/compare/v0.4.0...v0.5.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.4.0 - 2026-03-19

### 🚀 Enhancements

- add line number support in code blocks ([aef197d](https://github.com/indaco/herald/commit/aef197d)) ([#7](https://github.com/indaco/herald/pull/7))

### 🩹 Fixes

- theme colors ([5604a94](https://github.com/indaco/herald/commit/5604a94)) ([#8](https://github.com/indaco/herald/pull/8))

### 📖 Documentation

- add demos for built in themes ([cfdf21e](https://github.com/indaco/herald/commit/cfdf21e)) ([#9](https://github.com/indaco/herald/pull/9))

**Full Changelog:** [v0.3.0...v0.4.0](https://github.com/indaco/herald/compare/v0.3.0...v0.4.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.3.0 - 2026-03-19

### 🚀 Enhancements

- add GitHub-style alerts with Note, Tip, Important, Warning, Caution ([d00eb45](https://github.com/indaco/herald/commit/d00eb45)) ([#6](https://github.com/indaco/herald/pull/6))

### 🩹 Fixes

- use renderable Unicode icons for Note and Tip alerts ([9cb1491](https://github.com/indaco/herald/commit/9cb1491))

### 📖 Documentation

- update README with demo screenshots and restructured sections ([63ad497](https://github.com/indaco/herald/commit/63ad497))

### 🏡 Chores

- reorganize examples with demos subfolders ([3812aec](https://github.com/indaco/herald/commit/3812aec))
- add imagemagick to devbox ([ceb2100](https://github.com/indaco/herald/commit/ceb2100))
- add per-demo screenshot generation to justfile ([b4adbe3](https://github.com/indaco/herald/commit/b4adbe3))
- number-prefix example directories for logical ordering ([8efd5f7](https://github.com/indaco/herald/commit/8efd5f7))

**Full Changelog:** [v0.2.0...v0.3.0](https://github.com/indaco/herald/compare/v0.2.0...v0.3.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)

## v0.2.0 - 2026-03-18

### 🚀 Enhancements

- add nested list support with NestUL, NestOL, and hierarchical numbering ([86433a8](https://github.com/indaco/herald/commit/86433a8)) ([#5](https://github.com/indaco/herald/pull/5))
- add ColorPalette, adaptive themes, and built-in huh-compatible themes ([5e9e463](https://github.com/indaco/herald/commit/5e9e463)) ([#3](https://github.com/indaco/herald/pull/3))
- add CodeFormatter callback for pluggable syntax highlighting ([0fa72b7](https://github.com/indaco/herald/commit/0fa72b7)) ([#2](https://github.com/indaco/herald/pull/2))

### 🏡 Chores

- clean up justfile for library usage ([7c89299](https://github.com/indaco/herald/commit/7c89299))
- add freeze to devbox Go tools ([5292c9e](https://github.com/indaco/herald/commit/5292c9e))

**Full Changelog:** [v0.1.0...v0.2.0](https://github.com/indaco/herald/compare/v0.1.0...v0.2.0)

### ❤️ Contributors

- [@indaco](https://github.com/indaco)
