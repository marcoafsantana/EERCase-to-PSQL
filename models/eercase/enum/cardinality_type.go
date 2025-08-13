package enum

// CardinalityType – usando uint para enum
type CardinalityType uint

const (
	One CardinalityType = iota
	Many
)
