package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"multiplicator/internal/config"
	"multiplicator/internal/models"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type response struct {
	Result float64 `json:"result"`
}

type rtpParameter struct {
	rtp float64
}

var cfg *models.Config

func (s *rtpParameter) Sample() float64 {
	u := rand.Float64() // U ~ Uniform[0,1) [21]
	rtp := s.rtp
	if u < 1.0-rtp {
		return 1.0
	}
	if u >= 1.0-(rtp/10000.0) {
		return 10000.0
	}
	return rtp / (1.0 - u)
}

func (rtp *rtpParameter) MultiplierGenerator(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response{Result: rtp.Sample()}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// multiplicator - сервис создания мультипликаторов к некоторому набору клиентских параметроов
func main() {
	// parse rts parameter
	randNum := float64(rand.Intn(100)) + rand.Float64()
	rtpResp := flag.Float64("rtp", randNum, "rtp parameter")
	rtp := rtpParameter{rtp: *rtpResp}
	flag.Parse()

	// env Load
	cfg = config.LoadConfig()
	// add rtp to config
	// router init
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/get", rtp.MultiplierGenerator)

	log.Println("Server is run on port " + cfg.Port)
	if err := http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), router); err != nil {
		log.Fatal(err)
	}
}
