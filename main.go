package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"aidanwoods.dev/go-paseto"


)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Gagal memuat file .env")
		}
	} else {
		fmt.Println("File .env tidak ditemukan, pastikan file tersebut ada di direktori kerja.")
	}
}

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Gagal memuat file .env")
// 	}
// }

// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/HisyamSamAm
// @contact.email mhisyamnajwan@gmail.com

// @host localhost:8088
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	// Logging request
	app.Use(logger.New())

	// Basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Setup router
	router.SetupRoutes(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Endpoint not found",
		})
	})

	// Baca PORT dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port kalau tidak ada
	}

	log.Printf("Server is running at http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	fmt.Println("Private Key (hex):", privateKey.ExportHex())
	fmt.Println("Public Key (hex):", publicKey.ExportHex())
}