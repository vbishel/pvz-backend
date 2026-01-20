package repository

import (
	"orders-service/internal/config"
	"orders-service/internal/domain/product"
)

type productRepository struct {
	cfg *config.Config
}

func NewProductRepository(cfg *config.Config) product.ProductRepository {
	return &productRepository{cfg: cfg}
}

func (r *productRepository) Find(id product.ProductID) (product.Product, error) {
	panic("not implemented")
}

func (r *productRepository) FindByName(name string) (product.Product, error) {
	panic("not implemented")
}

func (r *productRepository) Create(name string) (product.ProductID, error) {
	panic("not implemented")
}

func (r *productRepository) Delete(id product.ProductID) error {
	panic("not implemented")
}
