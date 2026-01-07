package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (h *productHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	u, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	id := pgtype.UUID{Bytes: u, Valid: true}

	product, err := h.service.GetProductById(r.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.WriteJSON(w, http.StatusOK, map[string]interface{}{"product": product})
}
