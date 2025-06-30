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
)

func init() {

	if _, err := os.Stat(".env"); err != nil {
		if os.IsNotExist(err) {

			loadErr := godotenv.Load()
			if loadErr != nil {
				fmt.Println("Gagal memuat file .env:", loadErr)
			} else {
				fmt.Println("File .env berhasil dimuat.")
			}
		} else {

			fmt.Println("Terjadi error saat memeriksa file .env:", err)
		}
	}

}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Endpoint Tidak Ditemukan",
		})

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server berjalan di port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error Starting server: %v", err)
	}

}
