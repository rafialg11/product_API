package services_test

import (
	"errors"
	"product_api/entities"
	"product_api/mocks"
	"product_api/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepo)

	// Data dummy
	product := &entities.Product{
		Name:        "Test Product",
		Price:       100,
		Description: "Test Description",
		Quantity:    10,
	}

	t.Run("Success Save", func(t *testing.T) {
		mockRepo.EXPECT().Save(product).Return(product, nil)

		result, err := service.Save(product)
		assert.NoError(t, err)
		assert.Equal(t, product, result)
	})

	t.Run("Failed Save", func(t *testing.T) {
		mockRepo.EXPECT().Save(product).Return(nil, errors.New("failed to save product"))

		result, err := service.Save(product)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepo)

	t.Run("Success Find All", func(t *testing.T) {
		mockRepo.EXPECT().FindAll().Return([]entities.Product{}, nil)

		result, err := service.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed Find All", func(t *testing.T) {
		mockRepo.EXPECT().FindAll().Return(nil, errors.New("failed to fetch products"))

		result, err := service.FindAll()
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepo)
	var id uint = 1
	t.Run("Success Find By Id", func(t *testing.T) {
		mockRepo.EXPECT().FindById(id).Return(&entities.Product{}, nil)

		result, err := service.FindById(id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed Find By Id", func(t *testing.T) {
		mockRepo.EXPECT().FindById(id).Return(nil, errors.New("failed to fetch product"))

		result, err := service.FindById(id)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepo)

	product := &entities.Product{
		ID:          1,
		Name:        "Test Product",
		Price:       100,
		Description: "Test Description",
		Quantity:    10,
	}
	t.Run("Success Update", func(t *testing.T) {
		mockRepo.EXPECT().Update(product).Return(product, nil)
		result, err := service.Update(product)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed Update", func(t *testing.T) {
		mockRepo.EXPECT().Update(product).Return(nil, errors.New("failed to update product"))

		result, err := service.Update(product)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepo)

	var id uint = 1
	t.Run("Success Delete", func(t *testing.T) {
		mockRepo.EXPECT().Delete(id).Return(&entities.Product{}, nil)
		result, err := service.Delete(id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed Delete", func(t *testing.T) {
		mockRepo.EXPECT().Delete(id).Return(nil, errors.New("failed to delete product"))

		result, err := service.Delete(id)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
