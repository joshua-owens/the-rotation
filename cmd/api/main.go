package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joshua-owens/the-rotation/internal/scraper"
)

type config struct {
	port int
	env  string
}

type application struct {
	config  config
	logger  *log.Logger
	scraper *scraper.Scraper
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	scraper := scraper.New()

	app := &application{
		config:  cfg,
		logger:  logger,
		scraper: scraper,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %s", cfg.env, srv.Addr)
	srv.ListenAndServe()
}
