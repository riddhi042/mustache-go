package scanner

import "strings"

type Scanner struct {
	text string
	tail string
	pos  int
}

func New(text string) *Scanner {
	return &Scanner{
		text: text,
		tail: text,
		pos:  0,
	}
}

func (s *Scanner) EOS() bool {
	return s.tail == ""
}

func (s *Scanner) Scan(expected string) string {
	if strings.HasPrefix(s.tail, expected) {
		s.tail = s.tail[len(expected):]
		s.pos += len(expected)
		return expected
	}

	return ""
}

func (s *Scanner) ScanUntil(delimiter string) string {
	index := strings.Index(s.tail, delimiter)
	var match string

	switch index {
	case -1:
		match = s.tail
		s.tail = ""
	case 0:
		match = ""
	default:
		match = s.tail[:index]
		s.tail = s.tail[index:]
	}

	s.pos += len(match)

	return match
}
