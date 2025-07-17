// Application service layer
package services

import (
	"errors"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/Sup-Film/fiber-ecommerce-api/pkg/utils"
)

// AuthServiceImpl คือ struct ที่จะทำหน้าที่ implement ฟังก์ชั่นเกี่ยวกับ Auth ทั้งหมด
// โดยจำเป็นต้องมี userRepo (UserRepository) เพื่อใช้คุยกับฐานข้อมูล
type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewAuthService เป็น factory function ที่ใช้สร้าง instance ของ AuthServiceImpl
// วิธีนี้เรียกว่า Dependency Injection คือการ "ฉีด" dependency (userRepo) เข้ามา
func NewAuthService(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

// Register คือเมธอดสำหรับลงทะเบียนผู้ใช้ใหม่
func (s *AuthServiceImpl) Register(req *entities.RegisterRequest) (*entities.User, error) {
	// ตรวจสอบว่ามีอีเมลนี้ในระบบแล้วหรือยัง
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// เข้ารหัสผ่านก่อนบันทึกลงฐานข้อมูล
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// เตรียมข้อมูลผู้ใช้ใหม่
	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      entities.RoleUser, // กำหนดค่าเริ่มต้น Role เป็น User
		IsActive:  true,              // กำหนดค่าเริ่มต้นให้ Active
	}

	// เรียกใช้ Repository เพื่อสร้างผู้ใช้ในฐานข้อมูล
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login คือเมธอดสำหรับเข้าสู่ระบบ
func (s *AuthServiceImpl) Login(req *entities.LoginRequest) (*entities.LoginResponse, error) {
	// ค้นหาผู้ใช้ด้วยอีเมล
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// ตรวจสอบว่าผู้ใช้ถูกระงับหรือไม่
	if !user.IsActive {
		return nil, errors.New("user account is not active")
	}

	// ตรวจสอบรหัสผ่าน (ถ้าไม่ตรงกัน ให้ คืนค่า error)
	if !utils.CheckPassword(user.Password, req.Password) {
		return nil, errors.New("invalid email or password")
	}

	// สร้าง JWT Token สำหรับผู้ใช้
	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	// คืนค่าข้อมูลผู้ใช้พร้อมกับ Token
	return &entities.LoginResponse{
		User:  *user,
		Token: token,
	}, nil
}

// GetUserByID คือเมธอดสำหรับดึงข้อมูลผู้ใช้ตาม ID
func (s *AuthServiceImpl) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

// UpdateUser คือเมธอดสำหรับอัปเดตข้อมูลผู้ใช้
func (s *AuthServiceImpl) UpdateUser(user *entities.User) error {
	return s.userRepo.Update(user)
}
