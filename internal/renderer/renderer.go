package renderer

import (
	"fmt"
	"html"
	"strings"

	"github.com/riddhi042/mustache-go/internal/token"
)

func lookup(data map[string]any, name string) (any, bool) {
	parts := strings.Split(name, ".")

	var current any = data

	for _, part := range parts {
		currentMap, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}

		value, ok := currentMap[part]
		if !ok {
			return nil, false
		}

		current = value
	}

	return current, true
}

func Render(tokens []token.Token, data map[string]any) string {
	var out strings.Builder

	for _, t := range tokens {
		switch t.Type {

		case token.Text:
			out.WriteString(t.Value)

		case token.Name:
			if v, ok := lookup(data, t.Value); ok {
				escaped := html.EscapeString(fmt.Sprint(v))
				escaped = strings.ReplaceAll(escaped, "&#34;", "&quot;")
				escaped = strings.ReplaceAll(escaped, "/", "&#x2F;")
				out.WriteString(escaped)
			}

		case token.Unescaped:
			if v, ok := lookup(data, t.Value); ok {
				out.WriteString(fmt.Sprint(v))
			}

		case token.Section:
			if v, ok := lookup(data, t.Value); ok {
				if nestedData, ok := v.(map[string]any); ok {
					out.WriteString(Render(t.Children, nestedData))
				} else {
					out.WriteString(Render(t.Children, data))
				}
			}

		case token.InvertedSection:
			if _, ok := lookup(data, t.Value); !ok {
				out.WriteString(Render(t.Children, data))
			}

		case token.Comment:
			// Ignore comments
		}
	}

	return out.String()
}
