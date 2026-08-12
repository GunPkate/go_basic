package handlers

import (
	"sync"

	"github.com/google/uuid"

	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

// UserHandler holds dependencies for user-related routes.
// A mutex-guarded map stands in for a real database here.
type UserHandler struct {
	mu    sync.RWMutex
	users map[string]models.User
}

// NewUserHandler creates a UserHandler with an empty in-memory store.
func NewUserHandler() *UserHandler {
	return &UserHandler{
		users: make(map[string]models.User),
	}
}

// GetUsers returns all users.
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	list := make([]models.User, 0, len(h.users))
	for _, u := range h.users {
		list = append(list, u)
	}
	return c.JSON(list)
}

// GetUser returns a single user by ID.
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h.mu.RLock()
	user, ok := h.users[id]
	h.mu.RUnlock()

	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	return c.JSON(user)
}

// CreateUser adds a new user.
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	if input.Name == "" || input.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name and email are required",
		})
	}

	input.ID = uuid.NewString()

	h.mu.Lock()
	h.users[input.ID] = input
	h.mu.Unlock()

	return c.Status(fiber.StatusCreated).JSON(input)
}

// UpdateUser replaces an existing user's fields.
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h.mu.Lock()
	defer h.mu.Unlock()

	existing, ok := h.users[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if input.Name != "" {
		existing.Name = input.Name
	}
	if input.Email != "" {
		existing.Email = input.Email
	}
	h.users[id] = existing

	return c.JSON(existing)
}

// DeleteUser removes a user by ID.
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.users[id]; !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	delete(h.users, id)

	return c.SendStatus(fiber.StatusNoContent)
}
