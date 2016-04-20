package main

import (
	"bufio"
	"fmt"
	"os"
)

type TokenType int

const (
	TOctoHeadingLine TokenType = iota
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

	TIllegal
	TEof
)

var typeNames = map[TokenType]string{
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

	TIllegal: "TIllegal",
	TEof:     "TEof",
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
