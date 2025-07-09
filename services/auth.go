package services

import (
	"fmt"
	"go-project/database"
	"go-project/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key")

// Register user
func Register(username, password, email string) (*models.User, error) {
	// ตรวจสอบว่า username หรือ email ถูกใช้แล้วหรือยัง
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err == nil {
		return nil, fmt.Errorf("username already taken")
	}

	// การแฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// สร้างผู้ใช้ใหม่
	newUser := models.User{Username: username, Password: string(hashedPassword), Email: email}
	result := database.DB.Create(&newUser)
	return &newUser, result.Error
}

// Login user
func Login(username, password string) (string, error) {
	var user models.User
	// ตรวจสอบว่า username มีในฐานข้อมูลหรือไม่
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", fmt.Errorf("user not found")
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// สร้าง JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// เซ็นต์ JWT token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
