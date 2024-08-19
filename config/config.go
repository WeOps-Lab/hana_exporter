package config

// Config is the Go representation of the yaml config file.
type Config struct {
	Databases DatabaseConfig
}

// Credentials is the Go representation of the credentials section in the yaml
// config file.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Timeout  string
}
