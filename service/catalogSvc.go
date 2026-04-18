package service

import (
	"errors"
	dto_ "learning-backend/dto"
	"learning-backend/helper"
	"learning-backend/models"
	"learning-backend/repository"
)

var ErrProductNotFound = errors.New("product not found")

type  CatalogService struct {
	Auth helper.AuthHelper
  Repo repository.CatalogRepository
}



func (s CatalogService) CreateCategory(input dto_.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&models.Category{
		Name:         input.Name,
		ImageURL:     input.ImageURL,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatalogService) EditCategory(id int, input dto_.CreateCategoryRequest) (*models.Category, error) {

	exitCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")

	}

	if len(input.Name) > 0 {
		exitCat.Name = input.Name
	}

	if input.ParentID > 0 {
		exitCat.ParentID = input.ParentID
	}

	if len(input.ImageURL) > 0 {
		exitCat.ImageURL = input.ImageURL
	}

	if input.DisplayOrder > 0 {
		exitCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(exitCat)

	return updatedCat, err
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		// log the error
		return errors.New("category does not exist to delete")
	}

	return nil
}

func (s CatalogService) GetCategories() ([]*models.Category, error) {

	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("categories does not exist")
	}

	return categories, err

}

func (s CatalogService) GetCategory(id int) (*models.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")

	}
	return cat, nil
}

func (s CatalogService) CreateProduct(input dto_.CreateProductRequest, user models.User) (*models.Product, error) {

	shop, err := s.Repo.GetShopByUserID(user.ID)
  if err != nil {
    return nil, errors.New("Shop not found ")
  }

	
	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		ImageURL:    input.ImageURL,
		Slug:        helper.GenerateSlug(input.Name),
		ShopID:      uint(shop.ID),
		Stock:       uint(input.Stock),
	}

	err = s.Repo.CreateProduct(product)
	if err != nil {
		return nil, errors.New("failed to create product")
	}

	return product, nil
}



func (s CatalogService) EditProduct(id int, input dto_.CreateProductRequest, user models.User) (*models.Product, error) {

	exitProduct, err := s.Repo.FindProductById(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	// verify product owner
	if int(exitProduct.ShopID) != int(user.ID) {
		return nil, errors.New("you don't have manage rights of this product")
	}

	if len(input.Name) > 0 {
		exitProduct.Name = input.Name
	}

	if len(input.Description) > 0 {
		exitProduct.Description = input.Description
	}

	if input.Price > 0 {
		exitProduct.Price = input.Price
	}

	if input.CategoryID > 0 {
		exitProduct.CategoryID = input.CategoryID
	}

	updatedProduct, err := s.Repo.EditProduct(exitProduct)

	return updatedProduct, err
}

func (s CatalogService) DeleteProduct(id int, user models.User) error {
	exitProduct, err := s.Repo.FindProductById(id)
	if err != nil {
		return errors.New("product does not exist")
	}

	// verify product owner
	if int(exitProduct.ShopID) != int(user.ID) {
		return errors.New("you don't have manage rights of this product")
	}

	err = s.Repo.DeleteProduct(exitProduct)
	if err != nil {
		return errors.New("product cannot delete")
	}

	return nil
}



func (s CatalogService) GetProducts() ([]*models.Product, error) {
	products, err := s.Repo.FindProducts()
	if err != nil {
		return nil, errors.New("products does not exist")
	}

	return products, err
}

func (s CatalogService) GetProductById(id int) (*models.Product, error) {
	product, err := s.Repo.FindProductById(id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

func (s CatalogService) GetSellerProducts(id int) ([]*models.Product, error) {
	products, err := s.Repo.FindSellerProducts(id)
	if err != nil {
		return nil, errors.New("products does not exist")
	}

	return products, err
}

func (s CatalogService) UpdateProductStock(e models.Product) (*models.Product, error) {
	product, err := s.Repo.FindProductById(int(e.ID))
	if err != nil {
		return nil, errors.New("product not found")
	}

	// verify product owner
	if product.ShopID != e.ShopID {
		return nil, errors.New("you don't have manage rights of this product")
	}
	product.Stock = e.Stock
	editProduct, err := s.Repo.EditProduct(product)
	if err != nil {
		return nil, err
	}
	return editProduct, nil
}

func (s CatalogService) GetProductBySlug(shopSlug, productSlug string) (*models.Product, error) {

	// Step 1: find shop
	shop, err := s.Repo.FindShopBySlug(shopSlug)
	if err != nil {
		return nil, err
	}

	// Step 2: find product inside shop
	product, err := s.Repo.FindBySlugAndShop(productSlug, shop.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewCatalogService(auth helper.AuthHelper, repo repository.CatalogRepository) *CatalogService {
	return &CatalogService{
		Auth: auth,
		Repo: repo,
	}
}