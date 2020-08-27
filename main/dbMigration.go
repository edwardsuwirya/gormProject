package main

import (
	"github.com/edwardsuwirya/gormProject/entity"
)

func (a app) runMigration() {
	var migrator = a.db.Migrator()
	if migrator.HasTable("m_product") {
		migrator.DropTable("m_product")
	}
	if migrator.HasTable("m_category") {
		migrator.DropTable("m_category")
	}
	if err := a.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Category{}, &entity.Product{}); err != nil {
		panic(err)
	}

	products := []entity.Product{
		{
			ProductCode: "001",
			ProductName: "Cheeseburger",
			Category: entity.Category{
				CategoryName: "American Food",
			},
		},
		{
			ProductCode: "003",
			ProductName: "Cola Float",
			Category: entity.Category{
				CategoryName: "Drinks",
			},
		},
		{
			ProductCode: "004",
			ProductName: "French Fries",
			Category: entity.Category{
				CategoryName: "Appetizer",
			},
		},
		{
			ProductCode: "005",
			ProductName: "Nasi Goreng Kampung",
			Category: entity.Category{
				CategoryName: "Indonesian Food",
			},
		},
	}
	a.db.Create(products)

	cat := new(entity.Category)
	a.db.Where("category_name = ?", "American Food").First(cat)

	p := entity.Product{
		ProductCode: "002",
		ProductName: "Cesar Salad",
		CategoryId:  cat.ID,
	}
	a.db.Create(&p)
}
