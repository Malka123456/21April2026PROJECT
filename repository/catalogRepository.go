package repository

import (
	"errors"
	"learning-backend/models"

	"gorm.io/gorm"
	"log"
)

type CatalogRepository interface {
	CreateCategory(e *models.Category) error
	FindCategories() ([]*models.Category, error)
	FindCategoryById(id int) (*models.Category, error)
	EditCategory(e *models.Category) (*models.Category, error)
	DeleteCategory(id int) error

	CreateProduct(e *models.Product) error
	FindProducts() ([]*models.Product, error)
	FindProductById(id int) (*models.Product, error)
	FindSellerProducts(id int) ([]*models.Product, error)
	EditProduct(e *models.Product) (*models.Product, error)
	DeleteProduct(e *models.Product) error

	FindBySlugAndShop(slug string, shopID uint) (*models.Product, error)
	FindShopBySlug(slug string) (*models.Shop, error)
	GetShopByUserID(userID uint) (*models.Shop, error)
}

type catalogRepository struct {
	db *gorm.DB
}

func (c catalogRepository) CreateProduct(e *models.Product) error {
	err := c.db.Model(&models.Product{}).Create(e).Error
	if err != nil {
		log.Printf("err: %v", err)
		return errors.New("cannot create product")
	}
	return nil
}

func (c catalogRepository) FindProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := c.db.Preload("Shop").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c catalogRepository) FindProductById(id int) (*models.Product, error) {
	var product *models.Product
	err := c.db.First(&product, id).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("product does not exist")
	}
	return product, nil
}

func (c catalogRepository) FindSellerProducts(id int) ([]*models.Product, error) {
	var products []*models.Product
	err := c.db.Where("user_id=?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c catalogRepository) EditProduct(e *models.Product) (*models.Product, error) {
	err := c.db.Save(&e).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update product")
	}
	return e, nil
}

func (c catalogRepository) DeleteProduct(e *models.Product) error {
	err := c.db.Delete(&models.Product{}, e.ID).Error
	if err != nil {
		return errors.New("product cannot delete")
	}
	return nil
}

func (c catalogRepository) CreateCategory(e *models.Category) error {
	err := c.db.Create(&e).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("create category failed")
	}
	return nil
}

func (c catalogRepository) FindCategories() ([]*models.Category, error) {
	var categories []*models.Category

	err := c.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c catalogRepository) FindCategoryById(id int) (*models.Category, error) {
	var category *models.Category

	err := c.db.First(&category, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("category does not exist")
	}

	return category, nil
}

func (c catalogRepository) EditCategory(e *models.Category) (*models.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update category")
	}

	return e, nil
}

func (c catalogRepository) DeleteCategory(id int) error {

	err := c.db.Delete(&models.Category{}, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("fail to delete category")
	}

	return nil

}

func (r catalogRepository) FindBySlugAndShop(slug string, shopID uint) (*models.Product, error) {
	var product models.Product

	err := r.db.Where("slug = ? AND shop_id = ?", slug, shopID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r catalogRepository) FindShopBySlug(slug string) (*models.Shop, error) {
	var shop models.Shop

	err := r.db.Where("slug = ?", slug).First(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r catalogRepository) GetShopByUserID(userID uint) (*models.Shop, error) {
	var shop models.Shop

	err := r.db.Where("user_id = ?", userID).First(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}



func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}