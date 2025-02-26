package main

import (
	"EHR/database"
	"EHR/handler"
	"embed"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed public
var FS embed.FS

func main() {
	if err := database.InitializeDatabase(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))

	port := ":3000" 
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
