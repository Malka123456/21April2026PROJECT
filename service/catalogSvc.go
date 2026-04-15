package service

import (
	"errors"
	dto_ "learning-backend/dto"
	"learning-backend/helper"
	"learning-backend/models"
	"learning-backend/repository"
)


type  CatalogService struct {
	Auth helper.AuthHelper
  Repo repository.CatalogRepository
}



// package service

// import (
// 	"errors"
// 	"go-ecommerce-app/config"
// 	"go-ecommerce-app/internal/models"
// 	"go-ecommerce-app/internal/dto_"
// 	"go-ecommerce-app/internal/helper"
// 	"go-ecommerce-app/internal/repository"
// )

// type CatalogService struct {
// 	Repo   repository.CatalogRepository
// 	Auth   helper.Auth
// 	Config config.AppConfig
// }

func (s CatalogService) CreateCategory(input dto_.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&models.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
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

	if input.ParentId > 0 {
		exitCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		exitCat.ImageUrl = input.ImageUrl
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

func (s CatalogService) CreateProduct(input dto_.CreateProductRequest, user models.User) error {
	err := s.Repo.CreateProduct(&models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      int(user.ID),
		Stock:       uint(input.Stock),
	})

	return err
}

func (s CatalogService) EditProduct(id int, input dto_.CreateProductRequest, user models.User) (*models.Product, error) {

	exitProduct, err := s.Repo.FindProductById(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	// verify product owner
	if exitProduct.UserId != int(user.ID) {
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

	if input.CategoryId > 0 {
		exitProduct.CategoryId = input.CategoryId
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
	if exitProduct.UserId != int(user.ID) {
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
		return nil, errors.New("product does not exist")
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
	if product.UserId != e.UserId {
		return nil, errors.New("you don't have manage rights of this product")
	}
	product.Stock = e.Stock
	editProduct, err := s.Repo.EditProduct(product)
	if err != nil {
		return nil, err
	}
	return editProduct, nil
}

func NewCatalogService(auth helper.AuthHelper, repo repository.CatalogRepository) *CatalogService {
	return &CatalogService{
		Auth: auth,
		Repo: repo,
	}
}