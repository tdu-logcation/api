package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/tdu-logcation/api/handler"
	"github.com/tdu-logcation/api/utils"
)

var port string

func init() {
	env_port := os.Getenv("PORT")

	if len(env_port) == 0 {
		port = ":3000"
	} else {
		port = strings.Join([]string{":", env_port}, "")
	}
}

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", handler.RootHandler)

	corsConfig := utils.CorsConfig()
	handler := corsConfig.Handler(mux)

	if err := http.ListenAndServe(port, handler); err != nil {
		panic(err)
	}
}
