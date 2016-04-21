package main

import "fmt"

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok TokenType // last read token
		lit string    // last read literal
		n   int       // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser
func NewParser(s *Scanner) *Parser {
	return &Parser{
		s: s,
	}
}

func (p *Parser) scan() (tok TokenType, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreWhitespace() (tok TokenType, lit string) {
	tok, lit = p.scan()
	if tok == TWhitespace {
		tok, lit = p.scan()
	}
	return
}

// Schema represents the root level schema
type Schema struct {
	Tables []*CreateTable
}

// CreateTable represents a table as a whole
type CreateTable struct {
	IsPsuedo     bool
	TableName    string
	TableColumns []*TableColumn
}

type ColumnReferenceType int

const (
	ColumnRegular ColumnReferenceType = iota
	ColumnForeignKeyRegister
	ColumnForeignKeyReference
)

// TableColumn represents the column of a table
type TableColumn struct {
	ColumnName          string
	ColumnReferenceType ColumnReferenceType
}

func tokenArrayContains(t TokenType, types []TokenType) bool {
	for _, i := range types {
		if i == t {
			return true
		}
	}

	return false
}

// Parser parses scanner result
func (p *Parser) Parse() (*Schema, error) {
	sch := &Schema{}

	// First token should be a "SELECT" keyword.
	for {
		var tok TokenType
		var lit string
		if tok, lit = p.scanIgnoreWhitespace(); !tokenArrayContains(tok, HeadingTokens) {
			return nil, fmt.Errorf("found %q %s, expected Heading", lit, typeNames[tok])
		}

		tbl := &CreateTable{}
		if tok == TAtSignHeadingLine {
			tbl.IsPsuedo = true
		}

		if tok, lit = p.scanIgnoreWhitespace(); tok != TString {
			return nil, fmt.Errorf("found %q %s, expected TableName", lit, typeNames[tok])
		}

		tbl.TableName = lit
		sch.Tables = append(sch.Tables, tbl)

		break
	}

	return sch, nil

	// // Next we should loop over all our comma-delimited fields.
	// for {
	// 	// Read a field.
	// 	tok, lit := p.scanIgnoreWhitespace()
	// 	if tok != IDENT && tok != ASTERISK {
	// 		return nil, fmt.Errorf("found %q, expected field", lit)
	// 	}
	// 	stmt.Fields = append(stmt.Fields, lit)

	// 	// If the next token is not a comma then break the loop.
	// 	if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
	// 		p.unscan()
	// 		break
	// 	}
	// }

	// // Next we should see the "FROM" keyword.
	// if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
	// 	return nil, fmt.Errorf("found %q, expected FROM", lit)
	// }

	// // Finally we should read the table name.
	// tok, lit := p.scanIgnoreWhitespace()
	// if tok != IDENT {
	// 	return nil, fmt.Errorf("found %q, expected table name", lit)
	// }
	// stmt.TableName = lit

	// // Return the successfully parsed statement.
	// return stmt, nil
}
