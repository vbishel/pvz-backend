package config

import "time"

type (
	Config struct {
		App         `yaml:"app"`
		HTTP        `yaml:"http"`
		Log         `yaml:"logger"`
		Postgres    `yaml:"postgres"`
		AccessToken `yaml:"access_token"`
		CSRFToken   `yaml:"csrf_token"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port             string `env-required:"true" yaml:"port"               env:"HTTP_PORT"`
		CORSAllowOrigins string `env-required:"true" yaml:"cors_allow_origins" env:"CORS_ALLOW_ORIGINS"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Postgres struct {
		Host     string `env-required:"true" env:"PG_HOST"`
		Port     string `env-required:"true" env:"PG_PORT"`
		Database string `env-required:"true" env:"PG_DATABASE"`
		Username string `env-required:"true" env:"PG_USERNAME"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		Schema   string `env-required:"true" env:"PG_SCHEMA"`
	}

	AccessToken struct {
		TTL        time.Duration `env-required:"true" yaml:"ttl" env:"ACCESS_TOKEN_TTL"`
		SigningKey string        `env-required:"true" yaml:"signing_key" env:"ACCESS_TOKEN_SIGNING_KEY"`
	}

	CSRFToken struct {
		TTL       time.Duration `env-required:"true" yaml:"ttl"        env:"CSRF_TOKEN_TTL"`
		CookieKey string        `env-required:"true" yaml:"cookie_key" env:"CSRF_TOKEN_COOKIE_KEY"`
		HeaderKey string        `env-requried:"true" yaml:"header_key" env:"CSRF_TOKEN_HEADER_KEY"`
	}
)
