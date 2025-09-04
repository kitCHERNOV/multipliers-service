package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"multiplicator/internal/config"
	"net"
	"net/http"
)

// multiplicator - сервис создания мультипликаторов к некоторому набору клиентских параметроов
func main() {
	// env Load
	cfg := config.LoadConfig()
	// router init
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("get", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, {})
	})
	if _, err := net.Listen("tcp", cfg.Port); err != nil {
		log.Fatal(err)
	}
}
