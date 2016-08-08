package masq

import "fmt"

// Schema represents the root level schema
type Schema struct {
	Tables []*CreateTable
}

// CreateTable represents a table as a whole
type CreateTable struct {
	IsPsuedo     bool
	TableName    string
	TableComment []string
	TableColumns []*TableColumn

	PrimaryKey TableKeyColumns
	UniqueKeys map[int]TableKeyColumns
	Keys       map[int]TableKeyColumns

	// There can be zero or 1 AutoIncrement Column
	AutoIncrColumn *TableColumn
}

type TableKeyColumns []*TableColumn

// ColumnReferenceType is the parsed column reference type
//
//go:generate stringer -type=ColumnReferenceType
type ColumnReferenceType int

const (
	ColumnRegular ColumnReferenceType = iota
	ColumnForeignKeyRegister
	ColumnForeignKeyReference
)

// ColumnType represents the MySQL type of column
//
//go:generate stringer -type=ColumnType
type ColumnType int

const (
	_                        = iota
	ColumnTypeInt ColumnType = iota + 1
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

	ColumnTypeBool
)

type columnTypes map[ColumnType]string

func (c *columnTypes) getColumnType(input string) (ColumnType, error) {
	switch input {
	case "boolean":
		fallthrough
	case "bool":
		input = "tinyint"
	}

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
	ColumnComment       []string
	ColumnType          ColumnType
	ColumnSize          int
	ColumnReferenceType ColumnReferenceType

	DefaultValue *string

	Signed   bool
	Nullable bool

	AutoIncr bool
}
