package services

import (
	"context"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/google/uuid"
)

type paymentService struct {
	transactionRepo repositories.TransactionRepository
}

func NewPaymentService(transactionRepo repositories.TransactionRepository) services.PaymentService {
	return &paymentService{
		transactionRepo: transactionRepo,
	}
}

func (s *paymentService) CreatePayment(ctx context.Context, req *entities.CreatePaymentRequest) (*entities.Transaction, error) {
	return s.transactionRepo.Create(ctx, req)
}

func (s *paymentService) GetPaymentByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error) {
	return s.transactionRepo.GetByID(ctx, id)
}

func (s *paymentService) VerifyPayment(ctx context.Context, id uuid.UUID, req *entities.VerifyPaymentRequest) error {
	// ตรวจสอบข้อมูลการชำระเงิน
	transaction, err := s.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// ในการใช้งานจริง ควรมีการตรวจสอบกับ payment gateway
	// ที่นี่เป็นการจำลองการตรวจสอบ
	if transaction.TransactionID == req.TransactionID {
		// อัพเดทสถานะเป็น completed
		return s.transactionRepo.UpdateStatus(ctx, id, "completed")
	}

	// หากการตรวจสอบไม่ผ่าน อัพเดทสถานะเป็น failed
	return s.transactionRepo.UpdateStatus(ctx, id, "failed")
}

func (s *paymentService) CancelPayment(ctx context.Context, id uuid.UUID) error {
	return s.transactionRepo.Cancel(ctx, id)
}
