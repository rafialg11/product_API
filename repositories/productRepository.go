package repositories

import (
	"product_api/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product *entities.Product) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindById(id uint) (*entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(id uint) (*entities.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) Save(product *entities.Product) (*entities.Product, error) {
	return product, p.db.Create(product).Error
}

func (p *productRepository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	return products, p.db.Find(&products).Error
}

func (p *productRepository) FindById(id uint) (*entities.Product, error) {
	var product entities.Product
	return &product, p.db.First(&product, id).Error
}

func (p *productRepository) Update(product *entities.Product) (*entities.Product, error) {
	updatesProduct := &entities.Product{}
	if err := p.db.Model(&entities.Product{}).
		Where("id = ?", product.ID).
		Updates(product).
		First(updatesProduct, product.ID).Error; err != nil {
		return nil, err
	}
	return updatesProduct, nil
}

func (p *productRepository) Delete(id uint) (*entities.Product, error) {
	var product entities.Product
	return &product, p.db.Delete(&product, id).Error
}
