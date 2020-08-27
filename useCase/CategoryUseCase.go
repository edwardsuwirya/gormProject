package useCase

import (
	"github.com/edwardsuwirya/gormProject/entity"
	"github.com/edwardsuwirya/gormProject/repository"
)

type CategoryUseCase interface {
	RegisterCategory(category *entity.Category) *entity.Category
	GetCategoryCollection(withProduct bool) []*entity.Category
	GetCategoryById(id string) *entity.Category
	GetTotalCategory() *int64
}

type CategoryUseCaseImplementation struct {
	repo repository.CategoryRepository
}

func (p *CategoryUseCaseImplementation) RegisterCategory(category *entity.Category) *entity.Category {
	prod, err := p.repo.Create(category)
	if err != nil {
		return nil
	}
	return prod
}

func (p *CategoryUseCaseImplementation) GetCategoryCollection(withProduct bool) []*entity.Category {
	res, err := p.repo.FindAll(withProduct)
	if err != nil {
		return nil
	}
	return res
}

func (p *CategoryUseCaseImplementation) GetCategoryById(id string) *entity.Category {
	res, err := p.repo.Find(id)
	if err != nil {
		return nil
	}
	return res
}

func (p *CategoryUseCaseImplementation) GetTotalCategory() *int64 {
	res, err := p.repo.Count()
	if err != nil {
		return nil
	}
	return res
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImplementation{repo}
}
