package openapi

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	//MenuItem   MenuItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MenuItem   MenuItem
	MenuItemID int32
}
