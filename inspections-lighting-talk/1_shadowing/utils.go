package main

type Config struct {
	id int
}

func LoadConfig() *Config {
	return nil
}

func LoadConfigFromFile(path string) (*Config, error) {
	return &Config{}, nil
}
