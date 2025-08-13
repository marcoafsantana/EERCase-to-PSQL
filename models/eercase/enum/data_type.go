package enum

// DataType – usando uint para enum
type DataType uint

const (
	STRING DataType = iota
	BOOLEAN
	TIMESTAMP
	FLOAT
	INTEGER
	CLOB
	BLOB
)
