package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FabioCaffarello/go-expert/clean-architecture/internal/entity"
	"github.com/FabioCaffarello/go-expert/clean-architecture/internal/usecase"
	"github.com/FabioCaffarello/go-expert/clean-architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var dto usecase.ListOrdersInputDTO

	queryValues := r.URL.Query()
	page := queryValues.Get("page")
	limit := queryValues.Get("limit")

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.Page = pageNumber
	dto.Limit = limitNumber

	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
