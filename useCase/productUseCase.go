package useCase

import (
	"github.com/edwardsuwirya/gormProject/entity"
	"github.com/edwardsuwirya/gormProject/repository"
)

type ProductUseCase interface {
	RegisterProduct(product *entity.Product) *entity.Product
	GetProductCollection() []*entity.Product
	GetProductById(id string) *entity.Product
	GetTotalProduct() *int
}

type ProductUseCaseImplementation struct {
	repo repository.ProductRepository
}

func (p *ProductUseCaseImplementation) RegisterProduct(product *entity.Product) *entity.Product {
	prod, err := p.repo.Create(product)
	if err != nil {
		return nil
	}
	return prod
}

func (p *ProductUseCaseImplementation) GetProductCollection() []*entity.Product {
	res, err := p.repo.FindAll()
	if err != nil {
		return nil
	}
	return res
}

func (p *ProductUseCaseImplementation) GetProductById(id string) *entity.Product {
	res, err := p.repo.Find(id)
	if err != nil {
		return nil
	}
	return res
}

func (p *ProductUseCaseImplementation) GetTotalProduct() *int {
	res, err := p.repo.Count()
	if err != nil {
		return nil
	}
	return res
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImplementation{repo}
}
