package main

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

var typeNames = map[TokenType]string{
	TIllegal: "TIllegal",
	TEof:     "TEof",

	TOctoHeadingLine:   "TOctoHeading",
	TAtSignHeadingLine: "TAtSignHeading",

	TExclaimLine:  "TExclaimLine",
	TQuestionLine: "TQuestionLine",
	TDashLine:     "TDashLine",
	TColonLine:    "TColonLine",

	TAstrisk: "TAstrisk",
	TSigned:  "TSigned",
	// TNullable: "TNullable",

	TWhitespace: "TWhitespace",
	TNewLine:    "TNewLine",
	TString:     "TString",
	TComment:    "TComment",
}

// HeadingTokens lists tokens that are propper headings
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
