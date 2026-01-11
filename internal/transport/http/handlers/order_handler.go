package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shindeshubhamm/go-ecomm/internal/service"
	"github.com/shindeshubhamm/go-ecomm/internal/transport/http/json"
)

type orderHandler struct {
	service service.OrderService
}

func NewOrderHandler(svc service.OrderService) *orderHandler {
	return &orderHandler{
		service: svc,
	}
}

func (h *orderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.ListOrders(r.Context())
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}

	json.WriteJSON(w, http.StatusOK, orders)
}

func (h *orderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	u, err := uuid.Parse(idStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "Failed to parse id")
		return
	}

	id := pgtype.UUID{Bytes: u, Valid: true}
	order, err := h.service.GetOrderById(r.Context(), id)

	if err != nil {
		if err.Error() == "not found" {
			json.WriteError(w, http.StatusNotFound, "Order not found")
			return
		}
		json.WriteError(w, http.StatusInternalServerError, "Failed to fetch order")
		return
	}

	json.WriteJSON(w, http.StatusOK, order)
}
