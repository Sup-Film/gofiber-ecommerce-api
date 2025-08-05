package middleware

import (
	"strings"

	"github.com/Sup-Film/fiber-ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header เพื่อดึง Token
		authHeader := c.Get("Authorization")
		// ตรวจสอบว่า Authorization header มีค่าอยู่หรือไม่
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		// แยก Token ออกจาก Bearer
		tokenParts := strings.Split(authHeader, " ")
		// ตรวจสอบว่า Token มีรูปแบบที่ถูกต้องหรือไม่
		// ควรมี 2 ส่วน คือ "Bearer" และ Token
		// ถ้าไม่ถูกต้อง จะส่งกลับสถานะ 401 Unauthorized
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}
		token := tokenParts[1]

		// ตรวจสอบ Token ที่ได้รับ
		// โดยใช้ฟังก์ชัน ValidateJWT ที่เราได้สร้างไว้ใน utils
		// ถ้า Token ไม่ถูกต้อง จะคืนค่า error
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// ถ้า Token ถูกต้อง ให้เก็บข้อมูลผู้ใช้ใน context
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)

		// เรียกใช้ handler ถัดไปใน chain
		// เพื่อให้สามารถดำเนินการต่อได้
		return c.Next()
	}
}

// RoleMiddleware เป็น middleware สำหรับตรวจสอบบทบาทของผู้ใช้
func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ตรวจสอบว่า role ถูกเก็บไว้ใน context หรือไม่
		userRole := c.Locals("role")

		// ตรวจสอบว่า userRole ตรงกับบทบาทที่อนุญาตหรือไม่
		for _, role := range roles {
			if userRole == role {
				return c.Next() // ถ้าใช่ ให้ดำเนินการต่อ
			}
		}

		// ถ้าไม่ตรงกับบทบาทที่อนุญาต ส่งกลับสถานะ 403 Forbidden
		// เพื่อบอกว่าผู้ใช้ไม่มีสิทธิ์เข้าถึงทรัพยากรนี้
		// อาจจะเป็นเพราะบทบาทของผู้ใช้ไม่ตรงกับที่กำหนดไว้
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}
}
