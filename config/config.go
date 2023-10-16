package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	GetArangoDbURL string `yaml:"get_arangoDB_URL"`
	PostMySqlURL   string `yaml:"post_MySQL_URL"`
	GetSessionURL  string `yaml:"get_session_URL"`
	SqlLimit       int    `yaml:"sql_limit"`
	Login          string `env:"LOGIN"`
	Password       string `env:"PASSWORD"`
	Token          string `env:"TOKEN"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.UpdateEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
