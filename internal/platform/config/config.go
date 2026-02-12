package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/peterbourgon/ff/v4"
)

type Config struct {
	App      AppConfig
	HTTP     HTTPConfig
	Database DatabaseConfig
}

// Load loads the configuration from environment variables or a config file.
func Load(args []string) (*Config, error) {
	fs := flag.NewFlagSet("backend", flag.ContinueOnError)

	config := Config{}

	// Application settings
	fs.StringVar(&config.App.Env, "env", "development", "Application environment (development|production)")

	// HTTP server settings
	fs.StringVar(&config.HTTP.Port, "port", "8080", "HTTP server port")
	fs.DurationVar(&config.HTTP.ReadTimeout, "http-read-timeout", 5*time.Second, "HTTP server read timeout")
	fs.DurationVar(&config.HTTP.WriteTimeout, "http-write-timeout", 10*time.Second, "HTTP server write timeout")

	// Database settings
	fs.StringVar(&config.Database.Host, "db-host", "localhost", "Database host")
	fs.StringVar(&config.Database.Port, "db-port", "5432", "Database port")
	fs.StringVar(&config.Database.User, "db-user", "user", "Database user")
	fs.StringVar(&config.Database.Password, "db-password", "password", "Database password")
	fs.StringVar(&config.Database.Name, "db-name", "appdb", "Database name")
	fs.StringVar(&config.Database.SSLMode, "db-sslmode", "disable", "Database SSL mode (disable|require)")

	ff.Parse(fs, args,
		ff.WithEnvVars(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	return &config, nil
}

type AppConfig struct {
	Env string
}

type HTTPConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}
