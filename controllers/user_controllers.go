package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/fernode/goBookStore/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUserProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	claims := jwt.ExtractClaims(c)
	userID := uint(claims[identityKey].(float64))

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	claims := jwt.ExtractClaims(c)
	userID := uint(claims[identityKey].(float64))

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateInput models.User
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateInput.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updateInput.Password = string(hashedPassword)
	}

	if err := db.Model(&user).Updates(updateInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	claims := jwt.ExtractClaims(c)
	userID := uint(claims[identityKey].(float64))

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
