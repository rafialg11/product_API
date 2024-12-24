package repositories

import (
	"product_api/entities"

	"gorm.io/gorm"
)

type StockRepository interface {
	Save(stock *entities.Stock) (*entities.Stock, error)
	FindAll() ([]entities.Stock, error)
	FindById(id uint) (*entities.Stock, error)
	Update(stock *entities.Stock) (*entities.Stock, error)
	Delete(id uint) (*entities.Stock, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (s *stockRepository) Save(stock *entities.Stock) (*entities.Stock, error) {
	return stock, s.db.Create(stock).Error
}

func (s *stockRepository) FindAll() ([]entities.Stock, error) {
	var stocks []entities.Stock
	return stocks, s.db.Find(&stocks).Error
}

func (s *stockRepository) FindById(id uint) (*entities.Stock, error) {
	var stock entities.Stock
	return &stock, s.db.First(&stock, id).Error
}

func (s *stockRepository) Update(stock *entities.Stock) (*entities.Stock, error) {
	updatedStock := &entities.Stock{}
	if err := s.db.Model(&entities.Stock{}).
		Where("id = ?", stock.ID).
		Updates(stock).
		First(updatedStock, stock.ID).Error; err != nil {
		return nil, err
	}
	return updatedStock, nil
}

func (s *stockRepository) Delete(id uint) (*entities.Stock, error) {
	var stock entities.Stock
	return &stock, s.db.Delete(&stock, id).Error
}
