package repositories

import (
	"errors"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Createuser คือเมธอดสำหรับสร้างผู้ใช้ใหม่ในฐานข้อมูล
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	userModel := &models.User{}
	//
	userModel.FromEntity(user)

	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}

	*user = *userModel.ToEntity()
	return nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user models.User
	// .First(&user) ต้องส่ง address ของ user เพื่อให้ GORM สามารถแก้ไขข้อมูลได้
	// ถ้าไม่ใช้ address จะทำให้ GORM ไม่สามารถแก้ไขข้อมูลใน user ได้
	// และจะทำให้ไม่สามารถแปลงเป็น Entity ได้
	// ดังนั้นต้องใช้ pointer เพื่อให้ GORM สามารถแก้ไขข้อมูลได้
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // ถ้าไม่พบผู้ใช้ ให้คืนค่า nil
		}

		return nil, err
	}
	return user.ToEntity(), nil
}

func (r *UserRepositoryImpl) GetByID(id uint) (*entities.User, error) {
	var user models.User
	// .First(&user, id) ส่ง address ของ user เพื่อให้ GORM นำข้อมูลที่ได้มาใส่ใน user
	// id ตัวที่สองเป็นการระบุเงื่อนไขว่าให้ค้นหาจาก ID ที่ระบุ
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

func (r *UserRepositoryImpl) GetByRole(role entities.Role) ([]entities.User, error) {
	// ประกาศตัวแปร users เป็น slice ของ models.User
	var users []models.User

	// ใช้ GORM เพื่อค้นหาผู้ใช้ที่มี role ตรงกับที่ระบุ
	if err := r.db.Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}

	// ประกาศตัวแปร result เป็น slice ของ entities.User
	// เพื่อเก็บผลลัพธ์ที่แปลงจาก models.User เป็น entities.User
	var result []entities.User

	// วนลูปผ่านแต่ละ user ใน slice users
	// และแปลงแต่ละ user เป็น entities.User โดยใช้ ToEntity() เมธอด
	for _, user := range users {
		// ใช้ ToEntity() เพื่อแปลง models.User เป็น entities.User
		// แล้วเพิ่มเข้าไปใน result
		result = append(result, *user.ToEntity())
	}
	return result, nil
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)

	if err := r.db.Save(userModel).Error; err != nil {
		return err
	}

	*user = *userModel.ToEntity()
	return nil
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetAllUsers ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
func (r *UserRepositoryImpl) GetAll() ([]entities.User, error) {

	// ประกาศตัวแปร users เป็น slice ของ models.User
	// เพื่อเก็บข้อมูลผู้ใช้ที่ดึงมาจากฐานข้อมูล
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	// สร้างตัวแปรชื่อ result
	// เป็น slice (อาร์เรย์แบบขยายขนาดได้) ที่เก็บข้อมูลชนิด entities.User (คือ value ไม่ใช่ pointer)
	// ตอนเริ่มต้น slice นี้จะว่าง (ไม่มีสมาชิก)
	var result []entities.User

	// วนลูปผ่านแต่ละ user ใน slice users
	// และแปลงแต่ละ user เป็น entities.User โดยใช้ ToEntity() เมธอด
	// แล้วเพิ่มเข้าไปใน result
	for _, user := range users {
		result = append(result, *user.ToEntity())
	}
	return result, nil
}
