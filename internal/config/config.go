package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port           string
	AllowedOrigins []string
	Environment    string
	// MongodbUrl          string
	// MongodbPassword     string
	SupbaseUrl          string
	SupabaseAccessToken string
	SupabaseToken       string
	FrontEndUrl         string
}

func LoadConfig() (*Config, error) {
	env := strings.TrimSpace(getEnvWithDefaultValue("ENVIRONMENT", "development"))
	cfg := &Config{
		Port:        getEnvWithDefaultValue("PORT", "8080"),
		Environment: env,
		// MongodbUrl:          os.Getenv("MONGODB_URL"),
		// MongodbPassword:     os.Getenv("MONGODB_PASSWORD"),
		SupbaseUrl: os.Getenv("SUPABASE_URL"),
		// SupabaseAccessToken: os.Getenv("SUPABASE_ACCESS_TOKEN"),
		SupabaseToken: os.Getenv("SUPABASE_ANON_TOKEN"),
		FrontEndUrl:   getEnvWithDefaultValue("FRONTEND_URL", "http://localhost:8080"),
	}

	allowedOrigins := strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS"))
	if allowedOrigins == "" {
		if cfg.IsProduction() {
			return nil, fmt.Errorf("ALLOWED_ORIGINS is required in production")
		}
		origin := fmt.Sprintf("http://localhost:%s", cfg.Port)
		allowedOrigins = origin
	}

	cfg.AllowedOrigins = splitNTrim(allowedOrigins)

	// if cfg.MongodbPassword == "" {
	// 	return nil, fmt.Errorf("failed to load mongodb password")
	// }
	//
	// if cfg.MongodbUrl == "" {
	// 	return nil, fmt.Errorf("failed to load mongodb url")
	// }
	if cfg.SupabaseAccessToken == "" {
		return nil, fmt.Errorf("failed to load the supabase access token")
	}

	if cfg.SupabaseToken == "" {
		return nil, fmt.Errorf("failed to load the supabase token")
	}

	if cfg.SupbaseUrl == "" {
		return nil, fmt.Errorf("failed to load the supabase url")
	}

	if cfg.FrontEndUrl == "" {
		return nil, fmt.Errorf("frontend url is empty")
	}

	if cfg.Port == "" {
		return nil, fmt.Errorf("port provided is empty")
	}
	return cfg, nil
}

func getEnvWithDefaultValue(k, d string) string {
	if value := os.Getenv(k); value != "" {
		return value
	}
	return d
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func splitNTrim(input string) []string {
	parts := strings.Split(input, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
