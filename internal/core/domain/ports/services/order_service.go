package services

import (
	"context"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/google/uuid"
)

// OrderService interface สำหรับการจัดการคำสั่งซื้อ
type OrderService interface {
	CreateOrder(ctx context.Context, userID uuid.UUID, req *entities.CreateOrderRequest) (*entities.Order, error)
	GetOrders(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*entities.Order, error)
	CancelOrder(ctx context.Context, id uuid.UUID) error
	GetAllOrders(ctx context.Context, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateOrderStatusRequest) error
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, req *entities.UpdatePaymentStatusRequest) error
	UpdateShippingStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateShippingStatusRequest) error
}
