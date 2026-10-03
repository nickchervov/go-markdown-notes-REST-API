package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/adapters"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/controllers"
	"github.com/nickchervov/go-markdown-notes-REST-API/internal/service"
	"github.com/nickchervov/go-markdown-notes-REST-API/pkg/httpserver"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("loading environments: %v", err)
	}

	if err := RunApp(context.Background()); err != nil {
		log.Fatalf("running app: %v", err)
	}
}

func RunApp(ctx context.Context) error {
	repo, err := adapters.NewPostgres(ctx, os.Getenv("DSN"))
	if err != nil {
		return fmt.Errorf("creating repository: %w", err)
	}
	cache := adapters.NewRedis(os.Getenv("REDIS_HOST"))

	svc := service.New(repo, cache)
	routes := controllers.SetRoutes(svc)
	server := httpserver.New(os.Getenv("PORT"), routes)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("starting server: %v", err)
		}
	}()

	fmt.Printf("Server started at addr: %s\n", server.Addr)
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	fmt.Println("Starting graceful shutdown")

	shutdownCtx, stop := context.WithTimeout(ctx, 10*time.Second)
	defer stop()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}

	repo.Close()

	fmt.Println("Graceful shutdown completed")
	return nil
}
