package useCase

import (
	"github.com/edwardsuwirya/gormProject/entity"
	"github.com/edwardsuwirya/gormProject/repository"
)

type ProductUseCase interface {
	RegisterProduct(product *entity.Product) *entity.Product
	GetProductCollection(page, pageSize, order string) []*entity.Product
	GetProductById(id string) *entity.Product
	GetTotalProduct() *int64
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

func (p *ProductUseCaseImplementation) GetProductCollection(page, pageSize, order string) []*entity.Product {
	if len(order) == 0 {
		order = "id asc"
	}
	res, err := p.repo.FindAll(page, pageSize, order)
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

func (p *ProductUseCaseImplementation) GetTotalProduct() *int64 {
	res, err := p.repo.Count()
	if err != nil {
		return nil
	}
	return res
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImplementation{repo}
}
