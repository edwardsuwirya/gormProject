package repository

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/entity"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	FindAll() ([]*entity.Product, error)
	Find(id string) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	Count() (c *int, err error)
}

type ProductRepositoryImplementation struct {
	db *gorm.DB
}

func (p *ProductRepositoryImplementation) FindAll() ([]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryImplementation) Find(id string) (*entity.Product, error) {
	product := new(entity.Product)
	if err := p.db.Find(product, "id = ?", id).Error; err != nil {
		return nil, err
	}
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

func (p *ProductRepositoryImplementation) Count() (c *int,err error) {
	//err = p.db.Model(&entity.Product{}).Count(&c).Error
	err = p.db.Table("m_product").Count(&c).Error
	return
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImplementation{
		db,
	}
}
