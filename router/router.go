package router

import (
	"inibackend/handler"
	"github.com/gofiber/fiber/v2"
	"inibackend/config/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa",  handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", handler.GetAllMahasiswaByNPM)
	api.Post("/mahasiswa", middleware.Middlewares("admin"), handler.CreateMahasiswa)     // POST untuk tambah mahasiswa
	api.Put("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.UpdateMahasiswa)  // PUT untuk update mahasiswa
	api.Delete("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.DeleteMahasiswa)

	
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	
	

	 // DELETE untuk hapus mahasiswa
}
