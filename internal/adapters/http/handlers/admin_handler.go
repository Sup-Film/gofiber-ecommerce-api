package handlers

import (
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/Sup-Film/fiber-ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// ไม่ต้อง
type AdminHandler struct {
	authService services.AuthService
}

func NewAdminHandler(authService services.AuthService) *AdminHandler {
	return &AdminHandler{
		authService: authService,
	}
}

// GetAdminDashboard godoc
// @Summary Get admin dashboard
// @Description Get admin dashboard information (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/admin/dashboard [get]
func (h *AdminHandler) GetDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the admin dashboard",
		"user_id": c.Locals("userID"), // ใช้ userID ที่ได้จาก middleware
		"role":    c.Locals("role"),   // ใช้ role ที่ได้จาก middleware
	})
}

// AdminRegister godoc
// @Summary Register a new admin
// @Description Register a new admin user
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body entities.AdminRegisterRequest true "Admin registration data"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/admin/register [post]
func (h *AdminHandler) AdminRegister(c *fiber.Ctx) error {
	// ฟังก์ชันสำหรับการลงทะเบียนผู้ดูแลระบบ
	var req entities.AdminRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// ตรวจสอบข้อมูลที่ได้รับ
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed: " + err.Error(),
		})
	}

	user, err := h.authService.AdminRegister(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register admin: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
