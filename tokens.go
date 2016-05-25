package main

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
	TComment
)

// HeadingTokens lists tokens that are proper headings
var HeadingTokens = []TokenType{
	TOctoHeadingLine,
	TAtSignHeadingLine,
}

// ColumnTokens lists tokens that can begin column lines.
var ColumnTokens = []TokenType{
	TExclaimLine,
	TQuestionLine,
	TDashLine,
}

var TypeDecorators = []TokenType{
	TSigned,
	TAstrisk,
}
