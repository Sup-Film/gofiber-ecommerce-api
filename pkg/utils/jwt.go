package utils

import (
	"time"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims เป็น struct ที่กำหนดโครงสร้างข้อมูลที่จะถูกเก็บไว้ใน payload ของ JWT
// เราใช้ struct นี้เพื่อบอกว่าใน Token ของเราจะมีข้อมูลอะไรบ้าง
type Claims struct {
	UserID               uint   `json:"user_id"` // เก็บ User ID ของผู้ใช้
	Role                 string `json:"role"`    // เก็บ Role (บทบาท) ของผู้ใช้ เช่น 'admin' หรือ 'user'
	jwt.RegisteredClaims        // เป็นการฝัง (embed) struct RegisteredClaims จากไลบรารี jwt
	// ซึ่งจะช่วยให้เราสามารถใช้ Claims มาตรฐานของ JWT ได้ง่ายขึ้น
	// เช่น 'exp' (Expiration Time), 'iat' (Issued At), 'iss' (Issuer)
}

// GenerateJWT เป็นฟังก์ชั่นสำหรับสร้าง JWT Token ขึ้นมาใหม่
// รับค่า userID และ role ของผู้ใช้เป็นพารามิเตอร์ และจะคืนค่ากลับเป็น token (string) และ error (ถ้ามี)
func GenerateJWT(userID uint, role string) (string, error) {
	// ดึงค่า JWT_SECRET จากไฟล์ config ที่เราได้ตั้งค่าไว้
	// ค่านี้เป็น "ความลับ" ที่ใช้ในการเข้ารหัสและถอดรหัส Token
	secret := config.LoadConfig().JWTSecret

	// สร้าง claims หรือ payload สำหรับ Token นี้ โดยใส่ข้อมูล UserID, Role และกำหนดค่ามาตรฐานอื่นๆ
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt กำหนดเวลาหมดอายุของ Token (ในที่นี้คือ 24 ชั่วโมงนับจากเวลาปัจจุบัน)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// IssuedAt กำหนดเวลาที่ Token นี้ถูกสร้างขึ้น
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Issuer กำหนดชื่อของผู้ออก Token (ในที่นี้คือชื่อแอปพลิเคชันของเรา)
			Issuer: "fiber-ecommerce-api",
		},
	}

	// สร้าง Token ใหม่โดยใช้วิธีการเข้ารหัสแบบ HS256 (HMAC using SHA-256)
	// และใช้ claims ที่เราสร้างขึ้นเป็น payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ทำการ "เซ็น" (sign) Token ด้วย secret key ของเรา
	// ผลลัพธ์ที่ได้คือ string ของ Token ที่พร้อมใช้งาน
	// หากเกิดข้อผิดพลาดในขั้นตอนนี้ ฟังก์ชั่นจะคืนค่า error กลับไป
	return token.SignedString([]byte(secret))
}

// ValidateJWT เป็นฟังก์ชั่นสำหรับ "ตรวจสอบ" และ "ถอดรหัส" Token ที่ได้รับมา
// ฟังก์ชั่นจะรับ Token ที่เป็นข้อความ (string) เข้ามา
// ถ้า Token ถูกต้อง จะคืนค่าเป็นข้อมูล Claims (ข้อมูลที่เก็บใน Token) และไม่มี error
// ถ้า Token ไม่ถูกต้อง จะคืนค่าเป็น nil และ error
func ValidateJWT(tokenString string) (*Claims, error) {
	secret := config.LoadConfig().JWTSecret

	// jwt.ParseWithClaims คือหัวใจของการถอดรหัสและตรวจสอบ Token
	// มันจะพยายามแยกส่วนประกอบของ Token และเช็คลายเซ็นให้เรา
	// เราต้องใส่ 3 อย่างให้ฟังก์ชั่นนี้:
	// 1. tokenString: คือ Token ที่เราต้องการจะตรวจสอบ
	// 2. &Claims{}: คือ "พิมพ์เขียว" หรือโครงสร้างเปล่าๆ ของ Claims ของเรา
	//    ไลบรารีจะใช้พิมพ์เขียวนี้เพื่อลองแปลงข้อมูลใน Token กลับมาเป็น struct ของเรา
	// 3. ฟังก์ชั่นที่ทำหน้าที่คืน "กุญแจลับ" (secret key) กลับไป
	//    ไลบรารีจะเรียกใช้ฟังก์ชั่นนี้เพื่อขอกุญแจลับไปใช้ตรวจสอบลายเซ็นของ Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// เราแค่ส่งกุญแจลับของเรากลับไปในรูปแบบ []byte
		return []byte(secret), nil
	})

	// ตรวจสอบว่ามีข้อผิดพลาดเกิดขึ้นระหว่างการถอดรหัสหรือไม่
	// error ในขั้นตอนนี้อาจเกิดจากหลายสาเหตุ เช่น
	// - Token หมดอายุแล้ว (ไลบรารีเช็ค 'exp' claim ให้เราอัตโนมัติ)
	// - ลายเซ็น (Signature) ไม่ถูกต้อง
	// - รูปแบบของ Token ผิดเพี้ยนไป
	if err != nil {
		return nil, err
	}

	// หลังจากถอดรหัสแล้ว เราต้องทำการตรวจสอบ 2 อย่างสุดท้ายเพื่อให้แน่ใจจริงๆ
	// 1. `claims, ok := token.Claims.(*Claims)`:
	//    เป็นการเช็คว่าข้อมูลที่ถอดรหัสมาได้ (token.Claims) สามารถแปลงร่างเป็น struct `Claims` ของเราได้จริงๆ หรือไม่
	//    ถ้าแปลงได้ `ok` จะเป็น true
	// 2. `token.Valid`:
	//    เป็นการเช็คว่า Token นี้ "สมบูรณ์" หรือไม่ (เช่น ลายเซ็นถูกต้อง, ยังไม่หมดอายุ)
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// ถ้าทั้งสองอย่างผ่านฉลุย แสดงว่า Token นี้ถูกต้องและเป็นของเราจริง
		// เราก็คืนค่า claims ที่มีข้อมูลของผู้ใช้อยู่ข้างในกลับไป
		return claims, nil
	}

	// ถ้ามาถึงตรงนี้ได้ แสดงว่ามีบางอย่างผิดปกติ (เช่น Token ถูกต้องแต่ไม่ใช่ Claims รูปแบบที่เราต้องการ)
	// เราจะคืนค่า token เป็น nil และ error ว่าลายเซ็นไม่ถูกต้อง ซึ่งเป็น error ที่ปลอดภัยและครอบคลุม
	return nil, jwt.ErrSignatureInvalid
}
