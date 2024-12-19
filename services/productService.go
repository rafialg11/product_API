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
	product, err := p.repository.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) FindAll() ([]entities.Product, error) {
	products, err := p.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) FindById(id uint) (*entities.Product, error) {
	product, err := p.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) Update(product *entities.Product) (*entities.Product, error) {
	updatedProduct, err := p.repository.Update(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (p *productService) Delete(id uint) (*entities.Product, error) {
	delete, err := p.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	return delete, nil
}
