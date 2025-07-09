package middleware

import (
	"fmt"
	"go-project/database"
	"go-project/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("your-secret-key")

// Authenticate middleware สำหรับตรวจสอบ JWT token
func Authenticate(c *gin.Context) {
	// ดึง Authorization header จาก request
	tokenString := c.GetHeader("Authorization")

	// ตรวจสอบว่า header มี Authorization หรือไม่
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
		c.Abort() // หยุดการดำเนินการต่อ
		return
	}

	// ลบ "Bearer " ออกจาก token string
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// ตรวจสอบว่า token ถูกต้อง
	token, err := validateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// ตรวจสอบ claims ของ token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// ดึง user ID จาก claims
	userID := claims["sub"].(float64)

	// ดึงข้อมูลผู้ใช้จากฐานข้อมูล โดยใช้ userID ที่ได้จาก claims
	var user models.User
	if err := database.DB.Where("id = ?", uint(userID)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort() // หยุดการดำเนินการต่อถ้าผู้ใช้ไม่พบในฐานข้อมูล
		return
	}

	// เก็บข้อมูล user ใน context ของ Gin
	c.Set("user", user)

	c.Next() // ให้ดำเนินการกับ handler ถัดไป
}

// ฟังก์ชันที่ใช้ในการ validate token
func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	return token, err
}
