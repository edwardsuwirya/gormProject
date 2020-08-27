package main

import (
	"fmt"
	"github.com/edwardsuwirya/gormProject/entity"
)

func (a app) runMigration() {
	var migrator = a.db.Migrator()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%-30s %30s\n", "DB Migration", "[Failed]")
		}
	}()
	fmt.Printf("%-30s %30s\n", "Checking existing table", "[Done]")

	if migrator.HasTable("m_product") {
		if err := migrator.DropTable("m_product"); err != nil {
			fmt.Printf("%-30s %30s\n", "Dropping table product", "[Failed]")
		}
	}
	if migrator.HasTable("m_category") {
		if err := migrator.DropTable("m_category"); err != nil {
			fmt.Printf("%-30s %30s\n", "Dropping table category", "[Failed]")
		}
	}
	fmt.Printf("%-30s %30s\n", "Creating table structure", "[Done]")
	if err := a.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Category{}, &entity.Product{}); err != nil {
		fmt.Printf("%-30s %30s\n", "Creating table structure", "[Err]")
		panic(err)
	}

	fmt.Printf("%-30s %30s\n", "Insert some sample data", "[Done]")
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
	if err := a.db.Create(products).Error; err != nil {
		fmt.Printf("%-30s %30s\n", "Insert some sample data", "[Failed]")
		panic(err)
	}

	cat := new(entity.Category)
	a.db.Where("category_name = ?", "American Food").First(cat)

	p := entity.Product{
		ProductCode: "002",
		ProductName: "Cesar Salad",
		CategoryId:  cat.ID,
	}
	if err := a.db.Create(&p).Error; err != nil {
		fmt.Printf("%-30s %30s\n", "Insert some sample data", "[Failed]")
		panic(err)
	}
	fmt.Printf("%-30s %30s\n", "DB Migration", "[Done]")
}
