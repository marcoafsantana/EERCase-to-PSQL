package enum

// DisjointnessType – usando uint para enum
type DisjointnessType uint

const (
	OVERLAP DisjointnessType = iota
	DISJOINT
)
