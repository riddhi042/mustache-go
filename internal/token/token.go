package token

type TokenType int

const (
	Text TokenType = iota
	Variable
	Section
	EndSection
)

type Token struct {
	Type  TokenType
	Value string
}
