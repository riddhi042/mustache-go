package scanner

type Scanner struct {
	text string
	pos  int
}

func New(text string) *Scanner {
	return &Scanner{
		text: text,
		pos:  0,
	}
}

func (s *Scanner) EOS() bool {
	return s.pos >= len(s.text)
}
