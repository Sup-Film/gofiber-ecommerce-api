package utils

import (
	"errors"
	"regexp"
	"unicode"

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

// ValidatePasswordComplexity เป็นฟังก์ชั่นที่ใช้สำหรับตรวจสอบความซับซ้อนของรหัสผ่าน
func ValidatePasswordComplexity(password string) error {
	if len(password) < 8 {
		return errors.New("รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// ตรวจสอบว่ามีครบทุกประเภทหรือไม่
	if !hasUpper {
		return errors.New("รหัสผ่านต้องมีตัวอักษรใหญ่อย่างน้อย 1 ตัว")
	}
	if !hasLower {
		return errors.New("รหัสผ่านต้องมีตัวอักษรเล็กอย่างน้อย 1 ตัว")
	}
	if !hasNumber {
		return errors.New("รหัสผ่านต้องมีตัวเลขอย่างน้อย 1 ตัว")
	}
	if !hasSpecial {
		return errors.New("รหัสผ่านต้องมีอักขระพิเศษอย่างน้อย 1 ตัว (!@#$%^&*)")
	}

	return nil
}

// IsValidPassword เป็นฟังก์ชั่นที่ใช้สำหรับตรวจสอบรูปแบบของรหัสผ่าน
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	// ตรวจสอบแต่ละเงื่อนไขด้วย regex แยกกัน
	hasLower, _ := regexp.MatchString(`[a-z]`, password)
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	hasNumber, _ := regexp.MatchString(`\d`, password)
	hasSpecial, _ := regexp.MatchString(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~`+"`"+`]`, password)

	return hasLower && hasUpper && hasNumber && hasSpecial
}
