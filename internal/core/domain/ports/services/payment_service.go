package services

import (
	"context"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/google/uuid"
)

// PaymentService interface สำหรับการจัดการการชำระเงิน
type PaymentService interface {
	CreatePayment(ctx context.Context, req *entities.CreatePaymentRequest) (*entities.Transaction, error)
	GetPaymentByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error)
	VerifyPayment(ctx context.Context, id uuid.UUID, req *entities.VerifyPaymentRequest) error
	CancelPayment(ctx context.Context, id uuid.UUID) error
}
