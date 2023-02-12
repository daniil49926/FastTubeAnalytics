package api

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	PgDsn    string `toml:"postgres_dsn"`
}

func NewConfig() *Config {
	return &Config{}
}
