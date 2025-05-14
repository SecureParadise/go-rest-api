package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	// fmt.Printf("server started http://%s", cfg.HTTPServer.Addr)
	slog.Info("server started ", slog.String("URL", "http://"+cfg.HTTPServer.Addr))

	// graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("Shutting Down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
