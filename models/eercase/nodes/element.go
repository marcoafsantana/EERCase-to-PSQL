package nodes

// Element: extende Node
type Element struct {
	Node
	Name string `gorm:"type:varchar(255)" json:"name"`
}
