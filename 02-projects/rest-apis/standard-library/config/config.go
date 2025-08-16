package config

//  FIX: READ THE .env instead
type Config struct {
    PORT string
}

func LoadConfig() *Config {
    return &Config{ PORT: ":8080" }
}
