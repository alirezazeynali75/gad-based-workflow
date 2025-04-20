package configs

type Http struct {
	Port    string `env:"PORT" envDefault:"8081"`
	Address string `env:"ADDRESS" envDefault:"0.0.0.0"`
}