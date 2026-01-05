package service

import "context"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductService interface {
	ListProducts(ctx context.Context) ([]Product, error)
}

type svc struct {
	// db *sql.DB
}

func NewProductService() ProductService {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) ([]Product, error) {
	return []Product{
		{
			ID:   "1",
			Name: "Product 1",
		},
	}, nil
}
