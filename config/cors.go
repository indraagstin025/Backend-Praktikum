package config

var allowedOrigins = []string{
	"https://indrariksa.github.io/",
	"http://localhost:5173/",
	"https://frontend-praktikum-nu.vercel.app/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
