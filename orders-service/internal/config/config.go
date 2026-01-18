package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App         `yaml:"app"`
		HTTP        `yaml:"http"`
		Log         `yaml:"logger"`
		Postgres    `yaml:"postgres"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Env     string `env-required:"true" yaml:"app_env" env:"APP_ENV"`
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
		PoolMax  int    `env-required:"true" env:"PG_POOLMAX" yaml:"pool_max"`
	}
)

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string
	
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}