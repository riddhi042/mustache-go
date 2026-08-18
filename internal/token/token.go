package token

type TokenType int

const (
	Text TokenType = iota
	Name
	Section
	EndSection
	InvertedSection
	Partial
	Comment
	SetDelimiter
	Unescaped
)

type Token struct {
	Type     TokenType
	Value    string
	Start    int
	End      int
	Children []Token
}
