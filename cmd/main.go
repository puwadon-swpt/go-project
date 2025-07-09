package main

import (
	"go-project/database"
	"go-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// สร้าง Gin router
	r := gin.Default()

	// เชื่อมต่อกับฐานข้อมูล
	database.Connect()

	// ตั้งค่า routes
	routes.SetupRoutes(r)

	// รันเซิร์ฟเวอร์ที่ port 8080
	r.Run(":8080")
}
