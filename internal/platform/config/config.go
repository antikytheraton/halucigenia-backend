package config

import (
	"flag"
	"time"

	"github.com/peterbourgon/ff/v3"
	"github.com/pkazmier/ffyaml"
)

// Config holds the application configuration settings.
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
	fs.StringVar(&config.HTTP.Port, "port", "3000", "HTTP server port")
	fs.DurationVar(&config.HTTP.ReadTimeout, "http-read-timeout", 5*time.Second, "HTTP server read timeout")
	fs.DurationVar(&config.HTTP.WriteTimeout, "http-write-timeout", 10*time.Second, "HTTP server write timeout")

	// Database settings
	fs.StringVar(&config.Database.URL, "database-url", "", "Database connection URL")

	parser := ffyaml.New(
		ffyaml.WithKeyPath("services", "api"),
		ffyaml.WithAllowMissingKeyPath(true),
	)
	ff.Parse(fs, args,
		ff.WithEnvVars(),
		ff.WithConfigFile("config.yaml"),
		ff.WithConfigFileParser(parser.Parse),
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
	URL string
}
