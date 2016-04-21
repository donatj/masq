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
	TNullable

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

	TSigned:   "TSigned",
	TNullable: "TNullable",

	TWhitespace: "TWhitespace",
	TNewLine:    "TNewLine",
	TString:     "TString",
	TComment:    "TComment",

	TAstrisk: "TAstrisk",
}

// HeadingTokens lists tokens that are propper headings
var HeadingTokens = []TokenType{
	TOctoHeadingLine,
	TAtSignHeadingLine,
}
