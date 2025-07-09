package controllers

import (
	"go-project/database"
	"go-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API สำหรับดึงข้อมูลโปรไฟล์ของผู้ใช้
func GetProfile(c *gin.Context) {

	// ดึงข้อมูลผู้ใช้จาก context
	user := c.MustGet("user").(models.User)

	// ดึงข้อมูลจากฐานข้อมูล
	var profile models.User
	if err := database.DB.Where("id = ?", user.ID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         profile.ID,
		"username":   profile.Username,
		"email":      profile.Email,
		"created_at": profile.CreatedAt,
		"updated_at": profile.UpdatedAt,
	})
}
