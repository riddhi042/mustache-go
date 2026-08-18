package parser

import "github.com/riddhi042/mustache-go/internal/token"

func nestTokens(tokens []token.Token) []token.Token {
	var nested []token.Token
	collector := &nested

	var sections []*token.Token

	for _, tok := range tokens {
		switch tok.Type {

		case token.Section, token.InvertedSection:
			*collector = append(*collector, tok)

			current := &(*collector)[len(*collector)-1]

			sections = append(sections, current)

			collector = &current.Children

		case token.EndSection:
			if len(sections) > 0 {
				sections = sections[:len(sections)-1]
			}

			if len(sections) > 0 {
				collector = &sections[len(sections)-1].Children
			} else {
				collector = &nested
			}

		default:
			*collector = append(*collector, tok)
		}
	}

	return nested
}
