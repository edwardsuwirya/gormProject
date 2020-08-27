package repository

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, pageSize, order string) ([]*entity.Product, error)
	Find(id string) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	Count() (c *int64, err error)
}

type ProductRepositoryImplementation struct {
	db *gorm.DB
}

func (p *ProductRepositoryImplementation) FindAll(page, pageSize, order string) ([]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	if err := p.db.Order(order).Scopes(Paginate(page, pageSize)).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryImplementation) Find(id string) (*entity.Product, error) {
	product := new(entity.Product)
	//category := new(entity.Category)
	//product := new(entity.ProductWithCategoryDTO)
	//var col = []string{"m_product.id", "m_product.product_code", "m_product.product_name", "m_product.category_id", "m_category.category_name"}
	//if err := p.db.Model(product).Joins("JOIN m_category on m_category.id=m_product.category_id").Where("m_product.id = ?", id).Select(col).Scan(result).Error; err != nil {
	//	return nil, err
	//}
	//pretty.Println(result)
	//if err != nil {
	//	return nil, err
	//}
	//if err := p.db.Joins("JOIN m_category on m_category.id=m_product.category_id").Where("m_product.id = ?", id).Find(product).Error; err != nil {
	//	return nil, err
	//}
	//if err := p.db.Raw(`SELECT m_product.id,m_product.product_code,m_product.product_name,m_category.category_name
	//             FROM   m_product
	//              INNER JOIN m_category  on m_category.id=m_product.category_id
	//             WHERE m_product.id='` + id + "'").Scan(&product).Error; err != nil {
	//
	//}
	//var col = []string{"m_product.id", "m_product.product_code", "m_product.product_name", "m_product.category_id", "m_category.category_name"}
	//rows, err := p.db.Table("m_product").Joins("JOIN m_category on m_category.id=m_product.category_id").Where("m_product.id = ?", id).Select(col).Rows()
	p.db.Joins("Category").Where("m_product.id = ?", id).Find(product)
	//defer rows.Close()
	//for rows.Next() {
	//	if err = rows.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.CategoryId, &category.CategoryName); err != nil {
	//		config.AppConfig.Logger.Error(err)
	//		return nil, err
	//	}
	//	product.Category = *category
	//}
	//pretty.Println(product)
	return product, nil
}

func (p *ProductRepositoryImplementation) Create(product *entity.Product) (*entity.Product, error) {
	//tx := p.db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err := tx.Error; err != nil {
	//	return nil, err
	//}
	//if err := tx.Create(product).Error; err != nil {
	//	config.AppConfig.Logger.WithField("Database", "Create-Product").Error(err)
	//	tx.Rollback()
	//	return nil, err
	//}
	//
	//return tx.Value.(*entity.Product), tx.Commit().Error

	if err := p.db.Create(product).Error; err != nil {
		config.AppConfig.Logger.WithField("Database", "Create-Product").Error(err)
		return nil, err
	}
	return product, nil
}

func (p *ProductRepositoryImplementation) Count() (c *int64, err error) {
	//err = p.db.Model(&entity.Product{}).Count(&c).Error
	c = new(int64)
	err = p.db.Table("m_product").Count(c).Error
	return
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImplementation{
		db,
	}
}
