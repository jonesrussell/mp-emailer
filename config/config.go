package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonesrussell/loggo"
)

type Config struct {
	AppDebug       string
	AppEnv         string
	AppPort        string
	DBHost         string
	DBName         string
	DBPassword     string
	DBPort         string
	DBUser         string
	JWTExpiry      string
	JWTSecret      string
	LogFile        string
	LogLevel       string
	MailgunAPIKey  string
	MailgunDomain  string
	MailpitHost    string
	MailpitPort    string
	MigrationsPath string
	SessionName    string
	SessionSecret  string
}

// Load loads the configuration
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		AppDebug:       getEnv("APP_DEBUG", "false"),
		AppEnv:         getEnv("APP_ENV", "development"),
		AppPort:        getEnv("APP_PORT", "8080"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBName:         os.Getenv("DB_NAME"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBPort:         getEnv("DB_PORT", "3306"), // MariaDB default port
		DBUser:         os.Getenv("DB_USER"),
		JWTExpiry:      os.Getenv("JWT_EXPIRY"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		LogFile:        os.Getenv("LOG_FILE"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		MailgunAPIKey:  getEnv("MAILGUN_API_KEY", ""),
		MailgunDomain:  getEnv("MAILGUN_DOMAIN", ""),
		MailpitHost:    getEnv("MAILPIT_HOST", "localhost"),
		MailpitPort:    getEnv("MAILPIT_PORT", "1025"),
		MigrationsPath: getEnv("MIGRATIONS_PATH", "migrations"),
		SessionName:    getEnv("SESSION_NAME", "mpe"),
		SessionSecret:  os.Getenv("SESSION_SECRET"),
	}

	if config.LogFile == "" {
		config.LogFile = "mp-emailer.log" // Default value if not set
	}

	// Validate required variables
	if config.DBUser == "" || config.DBName == "" || config.DBPassword == "" {
		return nil, fmt.Errorf("DB_USER, DB_NAME, and DB_PASSWORD must be set in the environment")
	}
	if config.SessionSecret == "" {
		return nil, fmt.Errorf("SESSION_SECRET is not set in the environment")
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) GetLogLevel() loggo.Level {
	switch c.LogLevel {
	case "debug":
		return loggo.LevelDebug
	case "info":
		return loggo.LevelInfo
	case "warn":
		return loggo.LevelWarn
	case "error":
		return loggo.LevelError
	default:
		return loggo.LevelInfo
	}
}

func (c *Config) DatabaseDSN() string {
	// DSN format specific to MariaDB
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}
