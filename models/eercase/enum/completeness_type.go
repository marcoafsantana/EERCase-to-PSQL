package enum

// CompletenessType – usando uint para enum
type CompletenessType uint

const (
	PARTIAL CompletenessType = iota
	TOTAL
)
