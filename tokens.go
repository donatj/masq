package masq

// TokenType indicates the type of token being represented
//
//go:generate stringer -type=TokenType
type TokenType int

const (
	TIllegal TokenType = iota
	TEof

	TOctoHeadingLine
	TAtSignHeadingLine

	TExclaimLine
	TQuestionLine
	TDashLine
	TColonLine

	TAstrisk
	TSigned
	// TNullable

	TWhitespace
	TNewLine
	TString
	TEqualsString
	TComment
)

type TokenTypes []TokenType

// HeadingTokens lists tokens that are proper headings
var HeadingTokens = TokenTypes{
	TOctoHeadingLine,
	TAtSignHeadingLine,
}

// ColumnTokens lists tokens that can begin column lines.
var ColumnTokens = TokenTypes{
	TExclaimLine,
	TQuestionLine,
	TDashLine,
}

// TypeDecorators lists tokens that represent type decorations
var TypeDecorators = TokenTypes{
	TSigned,
	TAstrisk,
}

func (i TokenTypes) String() string {
	s := "["
	for x, j := range i {
		s += j.String()

		if x != len(i)-1 {
			s += ";"
		}
	}
	s += "]"

	return s
}
