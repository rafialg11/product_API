package repositories

import (
	"product_api/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(product *entities.Product) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindById(id uint) (*entities.Product, error)
	ChangePassword(id uint, password string) (*entities.Product, error)
	Delete(id uint) (*entities.Product, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(product *entities.Product) (*entities.Product, error) {
	return product, u.db.Create(product).Error
}

func (u *userRepository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	return products, u.db.Find(&products).Error
}

func (u *userRepository) ChangePassword(id uint, password string) (*entities.Product, error) {
	var product entities.Product
	return &product, u.db.Model(&product).Where("id = ?", id).Update("password", password).Error
}

func (u *userRepository) FindById(id uint) (*entities.Product, error) {
	var product entities.Product
	return &product, u.db.First(&product, id).Error
}

func (u *userRepository) Delete(id uint) (*entities.Product, error) {
	var product entities.Product
	return &product, u.db.Delete(&product, id).Error
}
