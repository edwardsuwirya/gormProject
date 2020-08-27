package repository

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(withProduct bool) ([]*entity.Category, error)
	Find(id string) (*entity.Category, error)
	Create(category *entity.Category) (*entity.Category, error)
	Count() (c *int64, err error)
}

type CategoryRepositoryImplementation struct {
	db *gorm.DB
}

func (p *CategoryRepositoryImplementation) FindAll(withProduct bool) ([]*entity.Category, error) {
	//categories := make([]*entity.Category, 0)
	//if err := p.db.Model(&entity.Category{}).Association("Products").Find(categories).Error; err != nil {
	//	return nil, err
	//}
	//return categories, nil
	cat := make([]*entity.Category, 0)
	//if err := p.db.Find(&cat).Error; err != nil {
	//	fmt.Printf(err.Error())
	//}
	//p.db.Model(&cat).Association("Products").Find(&cat.Products)

	//prods := make([]*entity.Product, 0)
	//p.db.Preload("Category").Find(&prods)
	//fmt.Printf("%v",prods)
	var err error
	if withProduct {
		err = p.db.Preload("Products").Find(&cat).Error
	} else {
		err = p.db.Find(&cat).Error
	}
	if err != nil {
		return nil, err
	}
	return cat, nil

}

func (p *CategoryRepositoryImplementation) Find(id string) (*entity.Category, error) {
	category := new(entity.Category)
	if err := p.db.Find(category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (p *CategoryRepositoryImplementation) Create(category *entity.Category) (*entity.Category, error) {
	if err := p.db.Create(category).Error; err != nil {
		config.AppConfig.Logger.WithField("Database", "Create-Category").Error(err)
		return nil, err
	}
	return category, nil
}

func (p *CategoryRepositoryImplementation) Count() (c *int64, err error) {
	err = p.db.Model(&entity.Category{}).Count(c).Error
	return
}

func NewCategoryRepo(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImplementation{
		db,
	}
}
