package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Login       string `gorm:"uniqueIndex"`
	Password    string
	IsAdmin     bool `gorm:"default:false"`
	IsStaff     bool `gorm:"default:false"`
	IsActive    bool `gorm:"default:true"`
	Cart        Cart
	CartID      uint
	BearerToken string `gorm:"index;default:null"`
}

type ShopItem struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"`
	Price    float64
	IsActive bool
}

type CartItem struct {
	gorm.Model
	ShopItem   ShopItem
	ShopItemId uint
	Cart       Cart
	CartID     uint
}

type Cart struct {
	gorm.Model
	Items []CartItem
}
