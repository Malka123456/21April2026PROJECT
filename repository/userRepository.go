package repository

import (
	"errors"
	"learning-backend/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(usr models.User) (models.User, error)
	FindUser(email string) (models.User, error)
	FindUserById(id uint) (models.User, error)
	UpdateUser(id uint, u models.User) (models.User, error)
	CreateBankAccount(e models.BankAccount) error

	// Cart
	FindCartItems(uId uint) ([]models.Cart, error)
	FindCartItem(uId uint, pId uint) (models.Cart, error)
	CreateCart(c models.Cart) error
	UpdateCart(c models.Cart) error
	DeleteCartById(id uint) error
	DeleteCartItems(uId uint) error

	// Order
	CreateOrder(o models.Order) error
	FindOrders(uId uint) ([]models.Order, error)
	FindOrderById(id uint, uId uint) (models.Order, error)

	// Profile
	CreateProfile(e models.Address) error
	UpdateProfile(e models.Address) error
   
	// Shop slug
	FindBySlug(slug string) (models.Shop, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateOrder(o models.Order) error {
	err := r.db.Create(&o).Error
	if err != nil {
		log.Printf("error on creating order %v", err)
		return errors.New("failed to create order")
	}
	return nil
}

func (r userRepository) FindOrders(uId uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id=?", uId).Find(&orders).Error
	if err != nil {
		log.Printf("error on fetching orders %v", err)
		return nil, errors.New("failed to fetch orders")
	}
	return orders, nil
}

func (r userRepository) FindOrderById(id uint, uId uint) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").Where("id=? AND user_id=?", id, uId).First(&order).Error
	if err != nil {
		log.Printf("error on fetching order %v", err)
		return models.Order{}, errors.New("failed to fetch order")
	}
	return order, nil
}

func (r userRepository) CreateProfile(e models.Address) error {
	err := r.db.Create(&e).Error
	if err != nil {
		log.Printf("error on creating profile with address %v", err)
		return errors.New("failed to create profile")
	}
	return nil
}

func (r userRepository) UpdateProfile(e models.Address) error {

	err := r.db.Where("user_id=?", e.UserID).Updates(e).Error
	if err != nil {
		log.Printf("error on update profile with address %v", err)
		return errors.New("failed to create profile")
	}
	return nil

}

func (r userRepository) FindCartItems(uId uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Where("user_id=?", uId).Find(&carts).Error
	return carts, err
}

func (r userRepository) FindCartItem(uId uint, pId uint) (models.Cart, error) {
	cartItem := models.Cart{}
	err := r.db.Where("user_id=? AND product_id=?", uId, pId).First(&cartItem).Error
	return cartItem, err
}

func (r userRepository) CreateCart(c models.Cart) error {
	return r.db.Create(&c).Error
}

func (r userRepository) UpdateCart(c models.Cart) error {
	var cart models.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.ID).Updates(c).Error
	return err
}

func (r userRepository) DeleteCartById(id uint) error {
	err := r.db.Delete(&models.Cart{}, id).Error
	return err
}

func (r userRepository) DeleteCartItems(uId uint) error {
	err := r.db.Where("user_id=?", uId).Delete(&models.Cart{}).Error
	return err
}

func (r userRepository) CreateBankAccount(e models.BankAccount) error {
	return r.db.Create(&e).Error
}

func (r userRepository) CreateUser(usr models.User) (models.User, error) {

	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("create user error %v", err)
		return models.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (r userRepository) FindUser(email string) (models.User, error) {

	var user models.User

	err := r.db.Preload("Address").First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return models.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (models.User, error) {

	var user models.User

	err := r.db.Preload("Address").
		Preload("Cart").
		Preload("Orders").
		First(&user, id).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return models.User{}, errors.New("user does not exist")
	}

	return user, nil
}
func (r userRepository) UpdateUser(id uint, u models.User) (models.User, error) {

	var user models.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("error on update %v", err)
		return models.User{}, errors.New("failed update user")
	}

	return user, nil
}

func (r userRepository) FindBySlug(slug string) (models.Shop, error) {
	var shop models.Shop
	err := r.db.Where("slug = ?", slug).First(&shop).Error
	return shop, err
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}