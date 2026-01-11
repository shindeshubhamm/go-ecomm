package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	repository "github.com/shindeshubhamm/go-ecomm/internal/adapters/postgresql/sqlc"
)

type OrderService interface {
	ListOrders(ctx context.Context) ([]repository.Order, error)
	GetOrderById(ctx context.Context, orderId pgtype.UUID) (repository.Order, error)
}

type orderSvc struct {
	repo repository.Querier
}

func NewOrderService(repo repository.Querier) OrderService {
	return &orderSvc{
		repo: repo,
	}
}

func (s *orderSvc) ListOrders(ctx context.Context) ([]repository.Order, error) {
	orders, err := s.repo.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderSvc) GetOrderById(ctx context.Context, orderId pgtype.UUID) (repository.Order, error) {
	order, err := s.repo.FindOrderById(ctx, orderId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repository.Order{}, errors.New("not found")
		}
		return repository.Order{}, err
	}

	return order, nil
}
