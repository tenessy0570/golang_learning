package database

import (
	"go-learn/internal"
	"go-learn/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB = GetDb()

func GetDb() gorm.DB {
	db, err := gorm.Open(sqlite.Open(internal.DbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return *db
}

func RecreateDb() {
	DropAllTables()
	CreateDbTables()
	FillDbWithTestRows()
}

func DropAllTables() {
	DB.Exec("drop table accounts")
	DB.Exec("drop table carts")
	DB.Exec("drop table cart_items")
	DB.Exec("drop table shop_items")
}
func CreateDbTables() {
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.CartItem{})
	DB.AutoMigrate(&models.ShopItem{})
}

func FillDbWithTestRows() {
	var cart1 = models.Cart{}
	var cart2 = models.Cart{}

	DB.Create(&cart1)
	DB.Create(&cart2)

	var user1 = models.Account{
		Login:    "tenessy",
		Password: "tenessy",
		IsAdmin:  true,
		IsStaff:  true,
		IsActive: true,
		CartID:   cart1.ID,
	}

	var user2 = models.Account{
		Login:    "tenessy1",
		Password: "tenessy1",
		IsAdmin:  false,
		IsStaff:  false,
		IsActive: true,
		CartID:   cart2.ID,
	}

	DB.Create(&user1)
	DB.Create(&user2)

	var shopItem1 = models.ShopItem{
		Name:     "Cocoa",
		Price:    5,
		IsActive: true,
	}
	var shopItem2 = models.ShopItem{
		Name:     "Milk",
		Price:    2,
		IsActive: true,
	}

	DB.Create(&shopItem1)
	DB.Create(&shopItem2)
}
