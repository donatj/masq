package masq

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

var eof = rune(0)

type Lexeme struct {
	Type  TokenType
	Value string
}

// Scanner represents a lexical scanner.
type Scanner struct {
	init bool
	r    *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *Scanner) peek() rune {
	r := s.read()
	s.unread()

	return r
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() Lexeme {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Lexeme{TWhitespace, buf.String()}
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	err := s.r.UnreadRune()
	if err != nil {
		panic(err)
	}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() Lexeme {
	// Read the next rune.
	ch := s.read()

	if isNewline(ch) || !s.init {
		s.init = true
		s.unread()
		return s.scanNewline()
	} else if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	if ch == '\'' || ch == '"' {
		s.unread()
		return s.scanString()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return Lexeme{TEof, ""}
	case '*':
		return Lexeme{TAstrisk, string(ch)}
	case '-':
		return Lexeme{TSigned, string(ch)}
	case '=':
		pch := s.peek()
		if pch == '\'' || pch == '"' {
			peq := s.scanString()
			peq.Type = TEqualsString

			return peq
		}

		peq := s.scanIdent()
		peq.Type = TEqualsString

		return peq
	}

	return Lexeme{TIllegal, string(ch)}
}

func (s *Scanner) scanNewline() Lexeme {
	var buf bytes.Buffer

	for {
		nl := s.read()
		buf.WriteRune(nl)

		if !isNewline(nl) && !isWhitespace(nl) {
			break
		}
	}

	s.unread()
	ch := s.read() // already in buffer

	switch ch {
	case eof:
		return Lexeme{TEof, buf.String()}
	case '@':
		return Lexeme{TAtSignHeadingLine, buf.String()}
	case '#':
		return Lexeme{TOctoHeadingLine, buf.String()}
	case '!':
		return Lexeme{TExclaimLine, buf.String()}
	case '?':
		return Lexeme{TQuestionLine, buf.String()}
	case ':':
		return s.scanColonLine()
	case '-':
		return Lexeme{TDashLine, buf.String()}
	}

	return Lexeme{TIllegal, buf.String()}
}

func (s *Scanner) scanString() Lexeme {
	var buf bytes.Buffer

	mark := s.read()
	prev := eof

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == mark && prev != '\\' {
			break
		} else {
			buf.WriteRune(ch)
			prev = ch
		}
	}

	return Lexeme{TString, buf.String()}
}

func (s *Scanner) scanColonLine() Lexeme {
	var buf bytes.Buffer

	if isWhitespace(s.peek()) {
		// this works... may just want to trim the result.
		s.scanWhitespace()
	}

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isNewline(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Lexeme{TColonLine, buf.String()}
}

func (s *Scanner) scanIdent() Lexeme {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return Lexeme{TString, buf.String()}
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	// unicode class n includes junk we don't want
	return (ch >= '0' && ch <= '9')
}

func isNewline(ch rune) bool {
	return ch == rune(10) || ch == rune(13)
}
