package types

import (
	"context"

	"github.com/Ibrahimss02/kitchen-oms/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
