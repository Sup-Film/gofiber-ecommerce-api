// ! Inbound port layer สำหรับการจัดการข้อมูลผู้ใช้
// Service port layer สำหรับการจัดการข้อมูลผู้ใช้
package services

import (
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
)

// AuthService interface กำหนดเมธอดที่ใช้ในการจัดการข้อมูลผู้ใช้ เช่น การลงทะเบียนผู้ใช้ใหม่, การเข้าสู่ระบบ, การดึงข้อมูลผู้ใช้ตาม ID และการอัปเดตข้อมูลผู้ใช้
type AuthService interface {
	Register(req entities.RegisterRequest) (*entities.User, error)
	AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error)
	Login(req entities.LoginRequest) (*entities.LoginResponse, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
