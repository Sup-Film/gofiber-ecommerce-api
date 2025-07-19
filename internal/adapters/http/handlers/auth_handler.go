// ไฟล์นี้ใช้สำหรับจัดการการยืนยันตัวตนของผู้ใช้
// รวมถึงการลงทะเบียนผู้ใช้ใหม่ การเข้าสู่ระบบ และการดึงข้อมูลโปรไฟล์ของผู้ใช้
// ใช้ Fiber framework สำหรับการจัดการ HTTP requests และ responses

package handlers

import (
	"strconv"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/Sup-Film/fiber-ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler สร้าง AuthHandler ใหม่
// รับพารามิเตอร์ authService ซึ่งเป็นบริการที่ใช้ในการจัดการการยืนยันตัวตน
// คืนค่า AuthHandler ที่พร้อมใช้งาน
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register
// Register godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.RegisterRequest true "Registration data"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req entities.RegisterRequest

	// ตรวจสอบว่า body ของ request สามารถถูกแปลงเป็น RegisterRequest ที่กำหนดไว้ได้หรือไม่
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบว่า request มีข้อมูลที่จำเป็นครบถ้วนหรือไม่
	// โดยใช้ utils.ValidateStruct เพื่อตรวจสอบความถูกต้องของข้อมูล
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// เรียกใช้ authService เพื่อทำการลงทะเบียนผู้ใช้ใหม่
	// ถ้ามีข้อผิดพลาดเกิดขึ้น จะส่งกลับสถานะ 409 Conflict
	user, err := h.authService.Register(req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// ส่งกลับสถานะ 201 Created พร้อมกับข้อมูลผู้ใช้ที่ลงทะเบียนสำเร็จ
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login ฟังก์ชันสำหรับเข้าสู่ระบบผู้ใช้
// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.LoginRequest true "Login credentials"
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entities.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.authService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

// GetUserProfile ฟังก์ชันสำหรับดึงข้อมูลโปรไฟล์ของผู้ใช้
// GetProfile godoc
// @Summary Get user profile
// @Description Get current user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entities.User
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/user/profile [get]
func (h *AuthHandler) GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.authService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}
