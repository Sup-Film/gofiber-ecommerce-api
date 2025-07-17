package repositories

import (
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
		return nil, err
	}
	return user.ToEntity(), nil
}

func (r *UserRepositoryImpl) GetById(id uint) (*entities.User, error) {
	var user models.User
	// .First(&user, id) ส่ง address ของ user เพื่อให้ GORM นำข้อมูลที่ได้มาใส่ใน user
	// id ตัวที่สองเป็นการระบุเงื่อนไขว่าให้ค้นหาจาก ID ที่ระบุ
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
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
