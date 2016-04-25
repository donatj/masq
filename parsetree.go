package main

import "fmt"

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

// ColumnReferenceType is the parsed column reference type
type ColumnReferenceType int

const (
	ColumnRegular ColumnReferenceType = iota
	ColumnForeignKeyRegister
	ColumnForeignKeyReference
)

// ColumnType represents the MySQL type of column
type ColumnType int

const (
	_                        = iota
	ColumnTypeInt ColumnType = iota
	ColumnTypeTinyInt
	ColumnTypeSmallInt
	ColumnTypeMediumInt
	ColumnTypeBigInt
	ColumnTypeFloat
	ColumnTypeDouble
	ColumnTypeDecimal
	ColumnTypeDate
	ColumnTypeDatetime
	ColumnTypeTimestamp
	ColumnTypeTime
	ColumnTypeYear
	ColumnTypeChar
	ColumnTypeVarchar
	// ColumnTypeBlob
	// ColumnTypeTinyBlob
	// ColumnTypeMediumBlob
	// ColumnTypeLongBlob
	ColumnTypeText
	ColumnTypeTinyText
	ColumnTypeMediumText
	ColumnTypeLongText
	// ColumnTypeEnum
)

type columnTypes map[ColumnType]string

func (c *columnTypes) getColumnType(input string) (ColumnType, error) {
	for t, c := range *c {
		if c == input {
			return t, nil
		}
	}

	return ColumnTypeInt, fmt.Errorf("unknown column type: %s", input)
}

var ColumnTypes = columnTypes{
	ColumnTypeInt:       "int",
	ColumnTypeTinyInt:   "tinyint",
	ColumnTypeSmallInt:  "smallint",
	ColumnTypeMediumInt: "mediumint",
	ColumnTypeBigInt:    "bigint",
	ColumnTypeFloat:     "float",
	ColumnTypeDouble:    "double",
	ColumnTypeDecimal:   "decimal",
	ColumnTypeDate:      "date",
	ColumnTypeDatetime:  "datetime",
	ColumnTypeTimestamp: "timestamp",
	ColumnTypeTime:      "time",
	ColumnTypeYear:      "year",
	ColumnTypeChar:      "char",
	ColumnTypeVarchar:   "varchar",
	// ColumnTypeBlob:       "blob",
	// ColumnTypeTinyBlob:   "tinyblob",
	// ColumnTypeMediumBlob: "mediumblob",
	// ColumnTypeLongBlob:   "longblob",
	ColumnTypeText:       "text",
	ColumnTypeTinyText:   "tinytext",
	ColumnTypeMediumText: "mediumtext",
	ColumnTypeLongText:   "longtext",
	// ColumnTypeEnum:       "enum",
}

// TableColumn represents the column of a table
type TableColumn struct {
	ColumnName          string
	ColumnType          ColumnType
	ColumnReferenceType ColumnReferenceType

	Signed   bool
	Nullable bool
}
