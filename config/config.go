package config

import (
	"fmt"
	"path/filepath"
	"time"
)

type Config struct {
	App          AppConfig      `yaml:"app"`
	Database     DatabaseConfig `yaml:"database" env:"sensitive"`
	Email        EmailConfig    `yaml:"email" env:"sensitive"`
	Auth         AuthConfig     `yaml:"auth" env:"sensitive"`
	Log          LogConfig      `yaml:"log"`
	Server       ServerConfig   `yaml:"server"`
	FeatureFlags FeatureFlags   `yaml:"feature_flags"`
}

type AppConfig struct {
	Debug bool        `env:"APP_DEBUG" envDefault:"false"`
	Env   Environment `env:"APP_ENV" envDefault:"development"`
	Host  string      `env:"APP_HOST" envDefault:"0.0.0.0"`
	Port  int         `env:"APP_PORT" envDefault:"8080"`
}

type DatabaseConfig struct {
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT" envDefault:"3306"`
	Name     string `env:"DB_NAME,required"`
}

type EmailConfig struct {
	Provider      EmailProvider `env:"EMAIL_PROVIDER" envDefault:"smtp"`
	MailgunAPIKey string        `env:"EMAIL_MAILGUN_API_KEY"`
	MailgunDomain string        `env:"EMAIL_MAILGUN_DOMAIN"`
	SMTP          SMTPConfig
}

type SMTPConfig struct {
	From     string `env:"EMAIL_FROM"`
	Host     string `env:"EMAIL_SMTP_HOST"`
	Password string `env:"EMAIL_SMTP_PASSWORD"`
	Port     int    `env:"EMAIL_SMTP_PORT" envDefault:"587"`
	Username string `env:"EMAIL_SMTP_USERNAME"`
}

type AuthConfig struct {
	JWTExpiry     string `env:"JWT_EXPIRY" envDefault:"24h"`
	JWTSecret     string `env:"JWT_SECRET,required"`
	SessionName   string `env:"SESSION_NAME" envDefault:"session"`
	SessionSecret string `env:"SESSION_SECRET,required"`
}

type LogConfig struct {
	File     string      `yaml:"file" env:"LOG_FILE" envDefault:"storage/logs/app.log"`
	Level    string      `yaml:"level" env:"LOG_LEVEL" envDefault:"info"`
	Format   string      `yaml:"format" env:"LOG_FORMAT" envDefault:"json"`
	Rotation LogRotation `yaml:"rotation"`
}

type ServerConfig struct {
	MigrationsPath              string
	RepresentativeLookupBaseURL string
}

type FeatureFlags struct {
	EnableMailgun bool `yaml:"enable_mailgun" env:"FEATURE_MAILGUN" envDefault:"false"`
	EnableSMTP    bool `yaml:"enable_smtp" env:"FEATURE_SMTP" envDefault:"true"`
	EnableMetrics bool `yaml:"enable_metrics" env:"FEATURE_METRICS" envDefault:"false"`
	BetaFeatures  bool `yaml:"beta_features" env:"FEATURE_BETA" envDefault:"false"`
}

type LogRotation struct {
	MaxSize    int  `yaml:"max_size" env:"LOG_MAX_SIZE" envDefault:"100"`
	MaxAge     int  `yaml:"max_age" env:"LOG_MAX_AGE" envDefault:"30"`
	MaxBackups int  `yaml:"max_backups" env:"LOG_MAX_BACKUPS" envDefault:"5"`
	Compress   bool `yaml:"compress" env:"LOG_COMPRESS" envDefault:"true"`
}

// DSN returns the database connection string
func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

// GetAbsolutePath returns the absolute path of a given path
func (c *Config) GetAbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return filepath.Clean(path)
	}
	return abs
}

// GetMigrationsPath returns the path to the migrations directory
func (c *Config) GetMigrationsPath() string {
	return c.Server.MigrationsPath
}

// GetLogFilePath returns the path to the log file
func (c *Config) GetLogFilePath() string {
	return c.Log.File
}

// GetJWTExpiryDuration returns the parsed JWT expiry duration
func (c *Config) GetJWTExpiryDuration() (time.Duration, error) {
	return time.ParseDuration(c.Auth.JWTExpiry)
}

func (c *Config) setupPaths() error {
	// Get absolute path for migrations
	migrationsPath, err := filepath.Abs(c.Server.MigrationsPath)
	if err != nil {
		return fmt.Errorf("invalid migrations path: %w", err)
	}
	c.Server.MigrationsPath = filepath.Clean(migrationsPath)

	// Get absolute path for log file
	logPath, err := filepath.Abs(c.Log.File)
	if err != nil {
		return fmt.Errorf("invalid log file path: %w", err)
	}
	c.Log.File = filepath.Clean(logPath)

	return nil
}
