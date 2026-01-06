package service

import (
	"context"

	repository "github.com/shindeshubhamm/go-ecomm/internal/adapters/postgresql/sqlc"
)

type ProductService interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
}

type svc struct {
	repo repository.Querier
}

func NewProductService(repo repository.Querier) ProductService {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repository.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
