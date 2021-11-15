package utils

import "github.com/rs/cors"

func CorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://log.tdu.app",
			"https://test-log.tdu.app",
			"https://local-log.tdu.app",
			"https://127.0.0.1:3000",
			"https://localhost:3000",
		},
		Debug: false,
	})
}
