package utils

import "github.com/rs/cors"

func CorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://log.tdu.app",
			"https://test-log.tdu.app",
			"https://local-log.tdu.app",
		},
		Debug: false,
	})
}
