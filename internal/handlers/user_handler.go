// api/handlers/user_handler.go
package handlers

import (
	"net/http"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    repo *repository.UserRepository
}

func NewUserHandler() *UserHandler {
    return &UserHandler{
        repo: repository.NewUserRepository(),
    }
}

// POST /api/v1/users
func (h *UserHandler) Create(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.repo.Create(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

// GET /api/v1/users
func (h *UserHandler) GetAll(c *gin.Context) {
    users, err := h.repo.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

// GET /api/v1/users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
    user, err := h.repo.FindByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// PUT /api/v1/users/:id
func (h *UserHandler) Update(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.repo.Update(c.Param("id"), &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DELETE /api/v1/users/:id
func (h *UserHandler) Delete(c *gin.Context) {
    if err := h.repo.Delete(c.Param("id")); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
