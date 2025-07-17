package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword เป็นฟังก์ชั่นที่ใช้สำหรับสร้างรหัสผ่านที่ถูกเข้ารหัส (hashed password)
func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword เป็นฟังก์ชั่นที่ใช้สำหรับตรวจสอบรหัสผ่านที่ผู้ใช้ป้อนเข้ามากับรหัสผ่านที่ถูกเข้ารหัส
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
