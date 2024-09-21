package handler

import (
	"net/http"

	"github.com/Ibrahimss02/kitchen-oms/services/common/genproto/orders"
	"github.com/Ibrahimss02/kitchen-oms/services/common/util"
	"github.com/Ibrahimss02/kitchen-oms/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewOrdersHttpHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest

	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	util.WriteJSON(w, http.StatusOK, res)
}
