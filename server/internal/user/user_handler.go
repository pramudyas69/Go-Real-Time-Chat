package user

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var u CreateUserReq
	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.Service.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var user LoginUserReq
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	u, err := h.Service.Login(c.Context(), &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    u.accesstoken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
		Domain:   "localhost",
	})

	return c.JSON(u)
}
