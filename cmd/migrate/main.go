package main

import (
	"flag"
	"log"
	"os"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/config"
)

func main() {
	var (
		up   = flag.Bool("up", false, "Run migrations")
		down = flag.Bool("down", false, "Rollback migrations (not implemented yet)")
	)
	flag.Parse()

	if !*up && !*down {
		log.Println("Usage:")
		log.Println("  go run cmd/migrate/main.go -up    # Run migrations")
		log.Println("  go run cmd/migrate/main.go -down  # Rollback migrations")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if *up {
		log.Println("Running database migrations...")
		err := config.RunMigrationManual(cfg)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migration completed successfully!")
	}

	if *down {
		log.Println("Rollback migrations not implemented yet")
		// TODO: Implement rollback functionality
	}
}
