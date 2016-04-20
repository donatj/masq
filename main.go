package main

import (
	"bufio"
	"fmt"
	"os"
)

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

type Token struct {
	Type TokenType
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s := NewScanner(r)

	for {
		tok, lit := s.Scan()
		fmt.Printf("%s: '%s'\n", typeNames[tok], lit)
		if tok == TEof {
			break
		}
	}

}
