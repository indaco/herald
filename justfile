# Variables

export LOG_STYLE := "emoji"
logger := "scripts/lib/logger.sh"

# Go commands

go := "go"
goclean := go + " clean"
freeze := "freeze"

# Default - show help
default:
    @just --list

# Clean the build directory and Go cache
clean:
    @. {{ logger }} && log_info "Clean the build directory and Go cache"
    rm -f coverage.out coverage.html
    {{ goclean }} -cache

# === Code Quality ===

# Format code
fmt:
    @. {{ logger }} && log_info "Running fmt"
    {{ go }} fmt ./...

# Run go-modernize with auto-fix
modernize:
    @. {{ logger }} && log_info "Running go-modernize"
    modernize --fix ./...

# Run golangci-lint
lint:
    @. {{ logger }} && log_info "Running golangci-lint"
    golangci-lint run ./...

# Run goreportcard-cli
reportcard:
    @. {{ logger }} && log_info "Running goreportcard-cli"
    goreportcard-cli -v

# Run govulncheck
security-scan:
    @. {{ logger }} && log_info "Running govulncheck"
    govulncheck ./...

# Run modernize, lint, and reportcard
check: fmt modernize lint reportcard

# Run go mod tidy
tidy:
    @. {{ logger }} && log_info "Running go mod tidy"
    {{ go }} mod tidy

# Run go mod download
deps:
    @. {{ logger }} && log_info "Running go mod download"
    {{ go }} mod download

# === Test Recipes ===

# Run all tests and print code coverage value
test:
    @. {{ logger }} && log_info "Run all tests"
    {{ go }} test $({{ go }} list ./... | grep -Ev 'examples/') -coverprofile=coverage.txt
    @. {{ logger }} && log_info "Total Coverage"
    {{ go }} tool cover -func=coverage.txt | grep total | awk '{print $3}'

# Clean go tests cache and run all tests
test-force:
    @. {{ logger }} && log_info "Clean go tests cache and run all tests"
    {{ go }} clean -testcache
    just test

# Run all tests and generate coverage report.
test-coverage:
    @. {{ logger }} && log_info "Run all tests and generate coverage report"
    {{ go }} test -count=1 -timeout 30s $({{ go }} list ./... | grep -Ev 'examples/') -covermode=atomic -coverprofile=coverage.txt

# Run all tests with race detector
test-race:
    @. {{ logger }} && log_info "Running tests with race detector"
    {{ go }} test -race $({{ go }} list ./... | grep -Ev 'examples/')

# === Utilities ===

# Update dependencies
deps-update:
    @. {{ logger }} && log_info "Running go update deps"
    {{ go }} get -u ./...
    {{ go }} mod tidy

# Capture and compose a single dark+light demo PNG
_capture-demo section:
    mkdir -p assets/demos
    HERALD_FORCE_DARK=1 {{ go }} run ./examples/demos/{{ section }}/ \
        | {{ freeze }} --output assets/demos/demo-{{ section }}-dark.png \
            --theme "Catppuccin Mocha" --padding 20 --window
    HERALD_FORCE_DARK=0 {{ go }} run ./examples/demos/{{ section }}/ \
        | {{ freeze }} --output assets/demos/demo-{{ section }}-light.png \
            --theme "Catppuccin Latte" --background "#FFFFFF" --padding 20 --window
    magick assets/demos/demo-{{ section }}-dark.png assets/demos/demo-{{ section }}-light.png \
        +append \( +clone -background black -shadow 60x20+0+10 \) \
        +swap -background none -layers merge +repage assets/demos/demo-{{ section }}.png
    rm -f assets/demos/demo-{{ section }}-dark.png assets/demos/demo-{{ section }}-light.png

# Capture and compose a single dark+light theme demo PNG
_capture-theme-demo name:
    mkdir -p assets/demos
    HERALD_FORCE_DARK=1 {{ go }} run ./examples/demos/builtin-themes/{{ name }}/ \
        | {{ freeze }} --output assets/demos/demo-theme-{{ name }}-dark.png \
            --theme "Catppuccin Mocha" --padding 20 --window
    HERALD_FORCE_DARK=0 {{ go }} run ./examples/demos/builtin-themes/{{ name }}/ \
        | {{ freeze }} --output assets/demos/demo-theme-{{ name }}-light.png \
            --theme "Catppuccin Latte" --background "#FFFFFF" --padding 20 --window
    magick assets/demos/demo-theme-{{ name }}-dark.png assets/demos/demo-theme-{{ name }}-light.png \
        +append \( +clone -background black -shadow 60x20+0+10 \) \
        +swap -background none -layers merge +repage assets/demos/demo-theme-{{ name }}.png
    rm -f assets/demos/demo-theme-{{ name }}-dark.png assets/demos/demo-theme-{{ name }}-light.png

# Generate all demo screenshots
demo-screenshot:
    just _capture-demo hero
    just _capture-demo headings
    just _capture-demo blocks
    just _capture-demo lists
    just _capture-demo alerts
    just _capture-demo inline
    just _capture-theme-demo dracula
    just _capture-theme-demo catppuccin
    just _capture-theme-demo base16
    just _capture-theme-demo charm

