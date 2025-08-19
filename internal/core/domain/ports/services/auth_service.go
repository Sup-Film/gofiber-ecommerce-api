// ! Inbound port layer สำหรับการจัดการข้อมูลผู้ใช้
// Service port layer สำหรับการจัดการข้อมูลผู้ใช้
package services

import (
	"context"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/google/uuid"
)

// AuthService interface กำหนดเมธอดที่ใช้ในการจัดการข้อมูลผู้ใช้ เช่น การลงทะเบียนผู้ใช้ใหม่, การเข้าสู่ระบบ, การดึงข้อมูลผู้ใช้ตาม ID และการอัปเดตข้อมูลผู้ใช้
type AuthService interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	AdminRegister(ctx context.Context, req *entities.AdminRegisterRequest) (*entities.User, error)
	Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
	RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error)
	Logout(ctx context.Context, userID uuid.UUID) error
	ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePasswordRequest) error
	ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error
	ValidateToken(ctx context.Context, token string) (*entities.User, error)
}
