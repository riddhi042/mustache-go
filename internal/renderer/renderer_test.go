package renderer

import (
	"testing"

	"github.com/riddhi042/mustache-go/internal/parser"
)

func TestVariable(t *testing.T) {
	template := "Hello {{name}}"

	tokens := parser.Parse(template)

	data := map[string]string{
		"name": "Riddhi",
	}

	got := Render(tokens, data)

	want := "Hello Riddhi"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
func TestMissingVariable(t *testing.T) {
	template := "Hello {{name}}"

	tokens := parser.Parse(template)

	got := Render(tokens, map[string]string{})

	want := "Hello "

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
func TestSection(t *testing.T) {
	template := "{{#user}}Hello{{/user}}"

	tokens := parser.Parse(template)

	got := Render(tokens, map[string]string{
		"user": "true",
	})

	want := "Hello"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestInvertedSection(t *testing.T) {
	template := "{{^user}}Hello{{/user}}"

	tokens := parser.Parse(template)

	got := Render(tokens, map[string]string{})

	want := "Hello"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
