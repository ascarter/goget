`goget` downloads and installs packages named by the import paths along with dependencies.

It is a drop-in replacement for `go get` and extends it to support vendoring in addition to the workspace. Using `-v` will fetch packages and dependencies to the project `vendor/` directory and write the vendor manifest.

# Install

		go get -u github.com/ascarter/goget

# Usage

Use `goget` whenever you would use `go get`. It passes through to `go get` for anything not specifically supported by `goget`. This approach is similar to other tools like `goimports`. The primary goal is to allow the get dependency resolution to be applied to vendored dependencies as supported by the `GO15VENDOREXPERIMENT` setting.

Flags:

		-v	vendor package

# Vendoring

Vendoring is the practice of copying third party dependencies directly into the project. Go 1.5 or later supports a `vendor/` directory with a layout identical to a Go workspace `src/` directory.
