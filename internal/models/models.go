package models

import "math/rand"

type Config struct {
	Server `yaml:"server" env-default:"true"`
	Amount
}

type Server struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost" env-required:"true"`
	Port string `yaml:"port" env:"PORT" env-default:"3000" env-required:"true"`
}

// Other models
type MultipliersResponse struct {
	Multipliers []float64
}

func NewMultiplicatorsResponse(amountNumbers int) *MultipliersResponse {
	multipliers := make([]float64, amountNumbers)
	rand.
	return &MultipliersResponse{Multipliers: make([]float64, amountNumbers)}
}
