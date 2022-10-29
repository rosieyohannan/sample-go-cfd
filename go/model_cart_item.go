package openapi

type CartItem struct {
	Id       int32
	MenuItem MenuItem `gorm:"embedded"`
}
