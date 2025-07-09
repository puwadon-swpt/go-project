package routes

import (
	"go-project/controllers"
	"go-project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes ตั้งค่า routes สำหรับแอปพลิเคชัน
func SetupRoutes(r *gin.Engine) {
	//TODO fix air does not work!
	// Route ที่ไม่ต้องการการยืนยันตัวตน
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Route ที่ต้องการการยืนยันตัวตน (middleware จะถูกใช้งานที่นี่)
	api := r.Group("/api")
	api.Use(middleware.Authenticate)            // ใช้ middleware สำหรับการตรวจสอบ JWT
	api.GET("/profile", controllers.GetProfile) // route สำหรับดึงข้อมูลโปรไฟล์ของผู้ใช้
}
