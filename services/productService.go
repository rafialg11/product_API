package services

import (
	"product_api/entities"
	"product_api/repositories"
)

type ProductService interface {
	Save(product *entities.Product) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindById(id uint) (*entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(id uint) (*entities.Product, error)
}

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (p *productService) Save(product *entities.Product) (*entities.Product, error) {
	return p.repository.Save(product)
}

func (p *productService) FindAll() ([]entities.Product, error) {
	return p.repository.FindAll()
}

func (p *productService) FindById(id uint) (*entities.Product, error) {
	return p.repository.FindById(id)
}

func (p *productService) Update(product *entities.Product) (*entities.Product, error) {
	updatedProduct, err := p.repository.Update(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (p *productService) Delete(id uint) (*entities.Product, error) {
	return p.repository.Delete(id)
}
