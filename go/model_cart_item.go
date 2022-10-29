package openapi

type CartItem struct {
	MenuItem MenuItem `gorm:"embedded"`
}
