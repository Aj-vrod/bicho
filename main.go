package main

import (
	"Aj-vrod/bicho/cmd"
	"Aj-vrod/bicho/internal/database"
	"Aj-vrod/bicho/pkg/config"
	"database/sql"
	"log"
)

var version = "dev-0.0.0"

func main() {
	// Connect to DB
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DBDNS)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db, "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connection and migrations successful")

	// Start the app
	rootCmd := cmd.NewRootCmd(version)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
