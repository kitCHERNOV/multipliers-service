package models

type Config struct {
	Server            `yaml:"server" env-default:"true"`
	MultipliersLimits NumberSet `yaml:"multipliers" env-required:"true"`
}

type NumberSet struct {
	LowerLimit float64 `yaml:"lower_limit"`
	UpperLimit float64 `yaml:"upper_limit"`
}

type Server struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost" env-required:"true"`
	Port string `yaml:"port" env:"PORT" env-default:"3000" env-required:"true"`
}
