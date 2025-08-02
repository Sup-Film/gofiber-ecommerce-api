// ! Outbound port layer สำหรับติดต่อกับฐานข้อมูลผ่าน
package repositories

import (
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
)

// UserRepository interface กำหนดเมธอดที่ใช้ในการจัดการข้อมูลผู้ใช้ ว่าจะให้ทำอะไรได้บ้าง
// เช่น การสร้างผู้ใช้ใหม่, ดึงข้อมูลผู้ใช้ตามอีเมล, ดึงข้อมูลผู้ใช้ตาม ID, อัปเดตข้อมูลผู้ใช้, ลบผู้ใช้, ดึงข้อมูลผู้ใช้ทั้งหมด และดึงข้อมูลผู้ใช้ตาม Role
type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByID(id uint) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	GetAll() ([]entities.User, error)
	GetByRole(role entities.Role) ([]entities.User, error)
}
