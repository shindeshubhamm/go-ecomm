package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repository "github.com/shindeshubhamm/go-ecomm/internal/adapters/postgresql/sqlc"
)

type ProductService interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
	GetProductById(ctx context.Context, id pgtype.UUID) (repository.Product, error)
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

func (s *svc) GetProductById(ctx context.Context, id pgtype.UUID) (repository.Product, error) {
	product, err := s.repo.FindProductById(ctx, id)
	if err != nil {
		return repository.Product{}, err
	}
	return product, nil
}
