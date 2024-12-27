package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	_ "embed"

	"github.com/hydrogen-1/goth-stats/internal/config"
	"github.com/hydrogen-1/goth-stats/internal/database"
	"github.com/hydrogen-1/goth-stats/internal/handlers"
	"github.com/hydrogen-1/goth-stats/internal/middleware"
	"github.com/hydrogen-1/goth-stats/internal/repository"
)

//go:embed sql/migrations/*.sql
var migrations embed.FS

func main() {
	context := context.Background()
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Error parsing config: %v\n", err)
	}

	pool, err := database.Connect(context, config, migrations)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()

	repository := repository.New(pool)

	statsHandler := handlers.StatsHandler{
		Repository: repository,
	}

	r := http.NewServeMux()

	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.HandleFunc("POST /visit", statsHandler.Visit)

	loggedMux := middleware.LoggingMiddleware(r)
	http.ListenAndServe("0.0.0.0:8000", loggedMux)
}
