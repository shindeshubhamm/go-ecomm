package handlers

import (
	"net/http"

	"github.com/shindeshubhamm/go-ecomm/internal/service"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(svc service.ProductService) *productHandler {
	return &productHandler{
		service: svc,
	}
}

func (h *productHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
}
