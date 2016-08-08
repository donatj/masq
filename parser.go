package main

import (
	"fmt"
	"log"
	"strings"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct { //todo replace with *Lexeme?
		lex Lexeme
		n   int // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser
func NewParser(s *Scanner) *Parser {
	return &Parser{
		s: s,
	}
}

func (p *Parser) scan() (lex Lexeme) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.lex
	}

	// Otherwise read the next token from the scanner.
	lex = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.lex = lex

	return
}

func (p *Parser) unscan() {
	//this is trash
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhitespace() (lex Lexeme) {
	lex = p.scan()
	if lex.Type == TWhitespace {
		lex = p.scan()
	}
	return
}

func tokenArrayContains(t TokenType, types []TokenType) bool {
	for _, i := range types {
		if i == t {
			return true
		}
	}

	return false
}

// Parse parses scanner result
func (p *Parser) Parse() (*Schema, error) {
	sch := &Schema{}

TableLoop:
	for {
		lex := p.scanIgnoreWhitespace()
		if lex.Type == TEof {
			break
		}

		if !tokenArrayContains(lex.Type, HeadingTokens) {
			return nil, fmt.Errorf("found %q %s, expected %s", lex.Value, lex.Type, HeadingTokens)
		}

		tbl := &CreateTable{
			UniqueKeys: make(map[int]TableKeyColumns),
			Keys:       make(map[int]TableKeyColumns),
		}
		if lex.Type == TAtSignHeadingLine {
			tbl.IsPsuedo = true
		}

		if lex = p.scanIgnoreWhitespace(); lex.Type != TString {
			return nil, fmt.Errorf("found %q %s, expected TableName", lex.Value, lex.Type)
		}

		tbl.TableName = lex.Value
		sch.Tables = append(sch.Tables, tbl)

	TblCommentLoop:
		for {
			comlex := p.scanIgnoreWhitespace()
			if comlex.Type == TColonLine {
				tbl.TableComment = append(tbl.TableComment, strings.TrimSpace(comlex.Value))
			} else {
				// I don't believe this is nessessary anymore
				// tbl.TableComment = strings.TrimSpace(tbl.TableComment)
				p.unscan()
				break TblCommentLoop
			}
		}

		// ColumnLoop:
		for {
			comlex := p.scanIgnoreWhitespace()
			if !tokenArrayContains(comlex.Type, ColumnTokens) {
				if tokenArrayContains(comlex.Type, HeadingTokens) {
					p.unscan()
					continue TableLoop
				}

				if comlex.Type == TEof {
					p.unscan()
					continue TableLoop
				}

				return nil, fmt.Errorf("found %q %s, expected %s", comlex.Value, comlex.Type, ColumnTokens)
			}

			col := &TableColumn{}

			switch comlex.Type {
			case TExclaimLine:
				col.ColumnReferenceType = ColumnForeignKeyRegister
			case TQuestionLine:
				col.ColumnReferenceType = ColumnForeignKeyReference
			case TDashLine:
				col.ColumnReferenceType = ColumnRegular
			default:
				return nil, fmt.Errorf("unexpected token: %s", comlex.Type)
			}

			colnlex := p.scanIgnoreWhitespace()
			if colnlex.Type != TString {
				return nil, fmt.Errorf("found %q %s, expected Column Name", colnlex.Value, colnlex.Type)
			}

			col.ColumnName = colnlex.Value

		ModLoop:
			for {
				modlex := p.scanIgnoreWhitespace()
				switch modlex.Type {
				case TAstrisk:
					col.Nullable = true
				case TSigned:
					col.Signed = true
				case TString:
					p.unscan()
					break ModLoop
				default:
					return nil, fmt.Errorf("unexpected token: %s", modlex.Type)
				}
			} //todo limit to one of each

			clex := p.scan()
			ctype, csize, err := strIntSuffixSplit(clex.Value)
			if err != nil {
				return nil, err
			}

			cctype, err := ColumnTypes.getColumnType(ctype)
			if err != nil {
				return nil, err
			}

			col.ColumnType = cctype
			col.ColumnSize = csize
			tbl.TableColumns = append(tbl.TableColumns, col)

			// KEYLOOP:
			for {
				kllex := p.scanIgnoreWhitespace()
				// log.Println(lit)
				log.Println("xxx", kllex.Type, kllex.Value)

				if kllex.Type != TString && kllex.Type != TAstrisk {
					log.Println("keyloop continuing", kllex.Type, kllex.Value)
					p.unscan()
					break
				}

				// autoIncr := false
				if kllex.Type == TAstrisk {
					// autoIncr = true
					if tbl.AutoIncrColumn == nil {
						tbl.AutoIncrColumn = col
					} else {
						return nil, fmt.Errorf("auto increment column already declared")
					}

					kllex = p.scanIgnoreWhitespace()
					if kllex.Type != TString {
						return nil, fmt.Errorf("found %q %s, expected %s", kllex.Value, kllex.Type, TString)
					}

					// continue KEYLOOP
				}

				log.Println("AAAAAAAAAAA", kllex)
				if kllex.Value == "pk" {
					log.Println("PRIMARY KEY", kllex)
					tbl.PrimaryKey = append(tbl.PrimaryKey, col)
				} else if kllex.Value[0:1] == "k" || kllex.Value[0:1] == "u" {
					log.Println("KEY", kllex)
					sPart, index, err := strIntSuffixSplit(kllex.Value)
					if err != nil || index <= 0 {
						return nil, fmt.Errorf("found '%s'; expected key name - %s", kllex.Value, err)
					}

					switch sPart {
					case "k":
						tbl.Keys[index] = append(tbl.Keys[index], col)
					case "u":
						tbl.UniqueKeys[index] = append(tbl.UniqueKeys[index], col)
					default:
						return nil, fmt.Errorf("found '%s'; expected key name - %s", kllex.Value, err)
					}

					// strIntSuffixSpkllex.Value
				} else {
					return nil, fmt.Errorf("found '%s'; expected key name - %s", kllex.Value, err)
				}
			}

		ColCommentLoop:
			for {
				comxlex := p.scanIgnoreWhitespace()
				if comxlex.Type == TColonLine {
					col.ColumnComment = append(col.ColumnComment, strings.TrimSpace(comxlex.Value))
				} else {
					// This shouldn't be nessessary
					// col.ColumnComment = strings.TrimSpace(col.ColumnComment)
					p.unscan()
					break ColCommentLoop
				}
			}
			// break
		}

		// break
	}

	return sch, nil
}
