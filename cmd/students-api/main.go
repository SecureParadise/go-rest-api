package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SecureParadise/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// sertup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	})
	// sertup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}
}
