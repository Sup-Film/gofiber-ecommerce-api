package entities

import (
	"time"
)

// กำหนดประเภทของ Role
type Role string

// กำหนดค่าคงที่สำหรับ Role
const (
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
)

// User struct แทนข้อมูลของผู้ใช้ในระบบ
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      Role      `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginRequest struct แทนข้อมูลที่ใช้ในการเข้าสู่ระบบ
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest struct แทนข้อมูลที่ใช้ในการลงทะเบียนผู้ใช้ใหม่
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type AdminRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      Role   `json:"role" validate:"required,oneof=admin user moderator"`
}

// LoginResponse struct แทนข้อมูลที่ส่งกลับหลังจากเข้าสู่ระบบสำเร็จ
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
