package renderer

import (
	"testing"

	"github.com/riddhi042/mustache-go/internal/parser"
)

func BenchmarkRender(b *testing.B) {
	template := "Hello {{name}}! Welcome to {{company}}."

	tokens := parser.Parse(template)

	data := map[string]string{
		"name":    "Riddhi",
		"company": "OpenAI",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Render(tokens, data)
	}
}
