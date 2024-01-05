package handlers

import (
    "journalapi/models"
    "net/http"
    "strconv"
	"log"

    "github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Add a new user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user   body    models.User   true  "Add User"
// @Success 200 {object} models.User
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := h.DB.Create(&user); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Summary Get a user
// @Description get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description update user's details by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "User ID"
// @Param   user   body    models.User   true  "Update User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// Add log here
	log.Printf("User after binding: %+v", user)


    // h.DB.Model(&user).Updates(user)
    h.DB.Save(&user)

    c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "User ID"
// @Success 204 "No Content"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User

    if result := h.DB.Delete(&user, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUsers godoc
// @Summary Get all users
// @Description get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
    var users []models.User
    if result := h.DB.Find(&users); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
        return
    }

    c.JSON(http.StatusOK, users)
}