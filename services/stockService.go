package services

import (
	"product_api/entities"
	"product_api/repositories"
)

type StockService interface {
	Save(stock *entities.Stock) (*entities.Stock, error)
	FindAll() ([]entities.Stock, error)
	FindById(id uint) (*entities.Stock, error)
	Update(stock *entities.Stock) (*entities.Stock, error)
	Delete(id uint) (*entities.Stock, error)
}

type stockService struct {
	repository repositories.StockRepository
}

func NewStockService(repository repositories.StockRepository) StockService {
	return &stockService{repository: repository}
}

func (s *stockService) Save(stock *entities.Stock) (*entities.Stock, error) {
	stock, err := s.repository.Save(stock)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func (s *stockService) FindAll() ([]entities.Stock, error) {
	stocks, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s *stockService) FindById(id uint) (*entities.Stock, error) {
	stock, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func (s *stockService) Update(stock *entities.Stock) (*entities.Stock, error) {
	stock, err := s.repository.Update(stock)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func (s *stockService) Delete(id uint) (*entities.Stock, error) {
	stock, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return stock, nil
}
