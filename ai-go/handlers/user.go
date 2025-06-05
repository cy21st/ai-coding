package handlers

import (
	"ai-go/database"
	"ai-go/middleware"
	"ai-go/models"
	"ai-go/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// Login handles user login
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	var user models.AdminUser
	if err := database.DB.Where("username = ? AND is_deleted = ?", req.Username, false).First(&user).Error; err != nil {
		utils.Unauthorized(c, "Invalid username or password", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Unauthorized(c, "Invalid username or password", nil)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		utils.InternalError(c, "Failed to generate token", err)
		return
	}

	utils.SuccessWithMessage(c, "Login successful", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalError(c, "Failed to hash password", err)
		return
	}

	user := models.AdminUser{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.InternalError(c, "Failed to create user", err)
		return
	}

	utils.SuccessWithMessage(c, "User created successfully", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

// GetUserInfo handles getting current user information
func GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("userID")
	var user models.AdminUser
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "User not found", err)
		return
	}

	utils.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

// GetUserList handles getting all users
func GetUserList(c *gin.Context) {
	var users []models.AdminUser
	if err := database.DB.Where("is_deleted = ?", false).Find(&users).Error; err != nil {
		utils.InternalError(c, "Failed to get user list", err)
		return
	}

	var userList []gin.H
	for _, user := range users {
		userList = append(userList, gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"role":       user.Role,
			"created_at": user.CreatedStr,
			"updated_at": user.UpdatedStr,
		})
	}

	utils.Success(c, userList)
}

// UpdateUser handles user update
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	var user models.AdminUser
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "User not found", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalError(c, "Failed to hash password", err)
		return
	}

	user.Username = req.Username
	user.Password = string(hashedPassword)
	user.Role = req.Role

	if err := database.DB.Save(&user).Error; err != nil {
		utils.InternalError(c, "Failed to update user", err)
		return
	}

	utils.SuccessWithMessage(c, "User updated successfully", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

// DeleteUser handles user deletion (soft delete)
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := database.DB.Model(&models.AdminUser{}).Where("id = ?", userID).Update("is_deleted", true).Error; err != nil {
		utils.InternalError(c, "Failed to delete user", err)
		return
	}

	utils.SuccessWithMessage(c, "User deleted successfully", nil)
}

// Logout handles user logout
func Logout(c *gin.Context) {
	utils.SuccessWithMessage(c, "Logout successful", nil)
}
