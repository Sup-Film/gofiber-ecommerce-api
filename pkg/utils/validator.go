package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// init ฟังก์ชันจะทำงานเมื่อ package ถูกโหลด
func init() {
	// ลงทะเบียน custom validator สำหรับรหัสผ่านที่ซับซ้อน
	validate.RegisterValidation("password_complex", validatePasswordComplex)
}

// validatePasswordComplex เป็นฟังก์ชันที่ใช้สำหรับตรวจสอบความซับซ้อนของรหัสผ่าน
// โดยฟังก์ชันนี้จะไปเรียกใช้ฟังก์ชัน IsValidPassword ที่อยู่ใน package utils
// และจะคืนค่า true ถ้ารหัสผ่านมีความซับซ้อนตามที่กำหนด
func validatePasswordComplex(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return IsValidPassword(password)
}

// ValidatePassword เป็นฟังก์ชันที่ใช้สำหรับตรวจสอบความถูกต้องของรหัสผ่าน
func ValidatePassword(password string) error {
	if !IsValidPassword(password) {
		if len(password) < 8 {
			return errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัวอักษร")
		}
		return errors.New("รหัสผ่านต้องมีตัวอักษรใหญ่ ตัวอักษรเล็ก ตัวเลข และอักขระพิเศษอย่างน้อยตัวละ 1 ตัว")
	}
	return nil
}
