package main

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

var eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
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
func (s *Scanner) scanWhitespace() (TokenType, string) {
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

	return TWhitespace, buf.String()
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok TokenType, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.

	// fmt.Println(ch)

	if isNewline(ch) {
		return TNewLine, string(ch)
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
		return TEof, ""
	case '@':
		return TAtSignHeadingLine, string(ch)
	case '#':
		return TOctoHeadingLine, string(ch)
	case '!':
		return TExclaimLine, string(ch)
	case ':':
		return s.scanColonLine()
	case '-':
		return TDashLine, string(ch)
	}

	return TIllegal, string(ch)
}

func (s *Scanner) scanString() (TokenType, string) {
	var buf bytes.Buffer

	mark := s.read()
	prev := eof

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == mark && prev != '\\' {
			break
		} else {
			_, _ = buf.WriteRune(ch)
			prev = ch
		}
	}

	return TString, buf.String()
}

func (s *Scanner) scanColonLine() (TokenType, string) {
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
			_, _ = buf.WriteRune(ch)
		}
	}

	return TColonLine, buf.String()
}

func (s *Scanner) scanIdent() (TokenType, string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	/*
		switch strings.ToUpper(buf.String()) {
		case "SELECT":
			return SELECT, buf.String()
		case "FROM":
			return FROM, buf.String()
		}
	*/

	// Otherwise return as a regular identifier.
	return TString, buf.String()
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
