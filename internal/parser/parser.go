package parser

import (
	"strings"

	"github.com/riddhi042/mustache-go/internal/scanner"
	"github.com/riddhi042/mustache-go/internal/token"
)

func Parse(template string) []token.Token {
	if template == "" {
		return []token.Token{}
	}

	s := scanner.New(template)

	var tokens []token.Token
	var sections []token.Token
	for !s.EOS() {
		// Read text before the next tag
		text := s.ScanUntil("{{")

		if text != "" {
			tokens = append(tokens, token.Token{
				Type:  token.Text,
				Value: text,
			})
		}

		// No more tags
		if s.EOS() {
			break
		}

		// Consume {{
		s.Scan("{{")

		triple := false

		// Check if this is {{{name}}}
		if s.Scan("{") != "" {
			triple = true
		}

		// Read tag contents
		var name string

		if triple {
			name = s.ScanUntil("}}}")
		} else {
			name = s.ScanUntil("}}")
		}

		if name == "" {
			s.Scan("}}")
			continue
		}

		tagType := token.Name

		switch name[0] {
		case '#':
			tagType = token.Section
			name = name[1:]
		case '/':
			tagType = token.EndSection
			name = name[1:]
		case '^':
			tagType = token.InvertedSection
			name = name[1:]
		case '!':
			tagType = token.Comment
			name = name[1:]
		case '>':
			tagType = token.Partial
			name = name[1:]
		case '&':
			tagType = token.Unescaped
			name = name[1:]
		}

		name = strings.TrimSpace(name)

		if triple {
			tagType = token.Unescaped
		}

		// Consume }}
		if triple {
			s.Scan("}}}")
		} else {
			s.Scan("}}")
		}
		tokens = append(tokens, token.Token{
			Type:  tagType,
			Value: name,
		})

		switch tagType {
		case token.Section, token.InvertedSection:
			sections = append(sections, token.Token{
				Type:  tagType,
				Value: name,
			})

		case token.EndSection:
			if len(sections) == 0 {
				panic("closing unopened section: " + name)
			}

			last := sections[len(sections)-1]

			if last.Value != name {
				panic("section mismatch: expected " + last.Value + ", got " + name)
			}

			sections = sections[:len(sections)-1]
		}
	}

	if len(sections) != 0 {
		last := sections[len(sections)-1]
		panic("unclosed section: " + last.Value)
	}

	return nestTokens(tokens)
}
