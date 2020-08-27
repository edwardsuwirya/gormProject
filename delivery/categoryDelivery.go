package delivery

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/repository"
	"github.com/edwardsuwirya/gormProject/useCase"
	"gorm.io/gorm"
)

type CategoryDelivery struct {
	categoryUseCase useCase.CategoryUseCase
}

func NewCategoryDelivery(db *gorm.DB) *CategoryDelivery {
	repo := repository.NewCategoryRepo(db)
	return &CategoryDelivery{
		categoryUseCase: useCase.NewCategoryUseCase(repo),
	}
}

func (cd *CategoryDelivery) Show() {
	res := cd.categoryUseCase.GetCategoryCollection(true)
	for _, p := range res {
		config.AppConfig.Logger.Info(p.ToString())
	}
}
