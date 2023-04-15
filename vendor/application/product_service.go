package application

import (
	"errors"

	"github.com/example/vendor/domain"
)

type ProductService interface {
	GetProduct(id string) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id string) error
}

type productService struct {
	productRepo domain.ProductRepository
}

func NewProductService(productRepo domain.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) GetProduct(id string) (*domain.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productService) GetAllProducts() ([]*domain.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) CreateProduct(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return s.productRepo.Save(product)
}

func (s *productService) UpdateProduct(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(id string) error {
	return s.productRepo.Delete(id)
}
