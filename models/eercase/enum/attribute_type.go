package enum

// AttributeType – usando uint para enum
type AttributeType uint

const (
	COMMON AttributeType = iota
	DERIVED
	MULTIVALUED
	IDENTIFIER
	DISCRIMINATOR
)
