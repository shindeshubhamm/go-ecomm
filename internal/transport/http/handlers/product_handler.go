package handlers

import (
	"log"
	"net/http"

	"github.com/shindeshubhamm/go-ecomm/internal/service"
	"github.com/shindeshubhamm/go-ecomm/internal/transport/http/json"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.WriteJSON(w, http.StatusOK, map[string]interface{}{"products": products})
}
