# mustache-go

A Go implementation of [Mustache](https://mustache.github.io/), the logic-less template syntax — ported from the reference [mustache.js](https://github.com/janl/mustache.js) implementation.

## Features

- Scanner, parser, and renderer built from scratch in Go
- Dot-notation variable lookup (e.g. `{{user.name.first}}`)
- Differential testing against the original `mustache.js` implementation to verify output parity
- Test coverage across the scanner, parser, and renderer packages

> **Status:** Core functionality is working. Some features and edge cases are still in progress — see [Roadmap](#roadmap) below.

## Project structure

```
mustache-go/
├── cmd/mustache/           # CLI entry point
├── internal/
│   ├── scanner/            # Tokenizes raw template text
│   ├── parser/             # Builds an AST from tokens
│   ├── renderer/           # Renders the AST against input data
│   └── token/              # Shared token definitions
├── tools/differential/     # Compares Go output against mustache.js (Node) for the same templates
├── mustache.js/            # Reference JS implementation, used for differential testing
└── go.mod
```

## Installation

```bash
go get github.com/riddhi042/mustache-go
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/riddhi042/mustache-go/internal/renderer"
)

func main() {
	data := map[string]any{
		"user": map[string]any{
			"name": "Riddhi",
		},
	}

	output, err := renderer.Render("Hello, {{user.name}}!", data)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	// Output: Hello, Riddhi!
}
```

> Note: adjust the example above to match your actual exported API — this is a starting point.

## Testing

Run the Go test suite:

```bash
go test ./...
```

### Differential testing

`tools/differential` renders the same templates through both this Go implementation and the original `mustache.js` (via Node) to catch behavioral differences:

```bash
cd tools/differential
npm install
go run main.go
```

## Roadmap

- [ ] Finish remaining Mustache spec compliance (partials, lambdas, etc. — update as applicable)
- [ ] Expand test coverage
- [ ] Polish CLI usage

## Acknowledgements

Built as a Go port of [mustache.js](https://github.com/janl/mustache.js) by the mustache.js contributors, used here as both a reference implementation and for differential testing.

## License

MIT License
