package main

import (
	"context"
	"fmt"
	R "golang-boilerplate/api/routes"
	"golang-boilerplate/ent/migrate"
	"golang-boilerplate/internal/config"
	"golang-boilerplate/internal/db"
	M "golang-boilerplate/internal/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	err := config.LoadEnvConfiguration(".", "config", "yaml")
	if err != nil {
		log.Fatalf("could not load configuration file: %w", err)
	}

	fmt.Println(config.Config.DB.Host)

	run()
}

func run() {

	app := fiber.New()
	M.SetupMiddleware(app)

	// DB Schema 마이그레이션
	setSchema()

	R.SetupRoutes(app)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	db.Client.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

func setSchema() {
	// Run the auto migration tool.
	if err := db.Client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
