package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_address"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
