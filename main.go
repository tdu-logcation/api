package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/tdu-logcation/api/handler"
	"github.com/tdu-logcation/api/utils"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	h2s := &http2.Server{}

	// Routes
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/user", handler.UserHandler)
	mux.HandleFunc("/log", handler.LogHandler)
	mux.HandleFunc("/rank", handler.RnakingHandler)

	corsConfig := utils.CorsConfig()
	handler := corsConfig.Handler(mux)

	server := &http.Server{
		Addr:    strings.Join([]string{"0.0.0.0", port}, ""),
		Handler: h2c.NewHandler(handler, h2s),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
