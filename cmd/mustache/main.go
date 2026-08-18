package main

import (
	"fmt"

	"github.com/riddhi042/mustache-go/internal/parser"
	"github.com/riddhi042/mustache-go/internal/renderer"
)

func main() {
	template := "Hello {{! this is ignored }}World"

	tokens := parser.Parse(template)

	data := map[string]string{
		"name": "<b>Riddhi</b>",
	}

	result := renderer.Render(tokens, data)

	fmt.Println(result)
}
