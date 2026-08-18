package parser

import "testing"

func TestEmptyTemplate(t *testing.T) {
	tokens := Parse("")

	if len(tokens) != 0 {
		t.Fatalf("expected 0 tokens, got %d", len(tokens))
	}
}

func TestTextOnly(t *testing.T) {
	tokens := Parse("Hello")

	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Value != "Hello" {
		t.Fatalf("expected %q, got %q", "Hello", tokens[0].Value)
	}
}

func TestVariable(t *testing.T) {
	tokens := Parse("{{name}}")

	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}

	if tokens[0].Value != "name" {
		t.Fatalf("expected %q, got %q", "name", tokens[0].Value)
	}
}

func TestSection(t *testing.T) {
	tokens := Parse("{{#user}}Hello{{/user}}")

	if len(tokens) == 0 {
		t.Fatal("expected tokens")
	}
}

func TestUnclosedSection(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for unclosed section")
		}
	}()

	Parse("{{#user}}")
}

func TestMismatchedSection(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for mismatched section")
		}
	}()

	Parse("{{#user}}{{/posts}}")
}
