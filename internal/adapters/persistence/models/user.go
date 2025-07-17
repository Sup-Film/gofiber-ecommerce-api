// package models ใช้เก็บ GORM models ที่เป็นตัวแทนของตารางในฐานข้อมูล

// NOTE:
// ในโปรเจกต์นี้มีการแยก Entity (domain/entities) กับ Model (adapters/persistence/models) อย่างชัดเจน
// - Entity คือโครงสร้างข้อมูลที่ใช้ใน business logic ของระบบ ไม่ผูกกับ database หรือ ORM
// - Model คือโครงสร้างที่ใช้กับ GORM เพื่อ map ข้อมูลกับฐานข้อมูลจริง มี tag สำหรับ DB และ JSON
// เวลาบันทึกหรืออ่านข้อมูลจากฐานข้อมูล จะต้องแปลงข้อมูลระหว่าง Entity กับ Model
// - ใช้ FromEntity เพื่อแปลง Entity → Model ก่อนบันทึกลง DB
// - ใช้ ToEntity เพื่อแปลง Model → Entity ก่อนนำไปใช้ใน business logic
// การแยกแบบนี้ช่วยให้โค้ดมีความยืดหยุ่นและแยกความรับผิดชอบตามหลัก Clean Architecture
package models

import (
	"time"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

// User แทน GORM model สำหรับตาราง 'users' ในฐานข้อมูล
// - กำหนดโครงสร้างและคุณสมบัติของแต่ละคอลัมน์
// - ใช้ GORM tags เพื่อระบุพฤติกรรม ORM และ json tags สำหรับ serialization
// - รองรับ Soft Delete และการจัดการเวลาโดยอัตโนมัติ
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`              // primaryKey: เป็น Primary Key ของตาราง
	Email     string         `gorm:"uniqueIndex;not null" json:"email"` // uniqueIndex: ค่าห้ามซ้ำ, not null: ห้ามเป็นค่าว่าง
	Password  string         `gorm:"not null" json:"-"`                 // json:"-": ไม่ต้องแสดงฟิลด์นี้ใน JSON ที่ส่งออกไป
	FirstName string         `gorm:"not null" json:"first_name"`
	LastName  string         `gorm:"not null" json:"last_name"`
	Role      entities.Role  `gorm:"type:varchar(20);default:'user'" json:"role"` // กำหนดชนิดข้อมูลและค่าเริ่มต้นใน DB
	IsActive  bool           `gorm:"default:true" json:"is_active"`               // กำหนดค่าเริ่มต้นเป็น true
	CreatedAt time.Time      `json:"created_at"`                                  // GORM จะจัดการเวลาสร้างให้เอง
	UpdatedAt time.Time      `json:"updated_at"`                                  // GORM จะจัดการเวลาอัปเดตให้เอง
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                              // สำหรับ Soft Delete, GORM จะจัดการให้
}

// ToEntity เป็นเมธอดที่ใช้แปลง GORM model (ข้อมูลจาก DB) ไปเป็น Domain Entity
// เพื่อให้ Service Layer นำไปใช้โดยไม่ต้องผูกติดกับโครงสร้างของฐานข้อมูล
func (u *User) ToEntity() *entities.User {
	return &entities.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromEntity เป็นฟังก์ชั่นที่ใช้แปลง Domain Entity ไปเป็น GORM model
// เพื่อเตรียมข้อมูลสำหรับบันทึกหรืออัปเดตลงในฐานข้อมูล
func (u *User) FromEntity(entity *entities.User) {
	u.ID = entity.ID
	u.Email = entity.Email
	u.Password = entity.Password
	u.FirstName = entity.FirstName
	u.LastName = entity.LastName
	u.Role = entity.Role
	u.IsActive = entity.IsActive
	u.CreatedAt = entity.CreatedAt
	u.UpdatedAt = entity.UpdatedAt
}
