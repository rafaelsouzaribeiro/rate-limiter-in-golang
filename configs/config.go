package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	HostRedis     string `env:"HOST_REDIS"`
	PasswordRedis string `env:"PASSWORD_REDIS"`
	MaxRequest    int    `env:"MAX_REQUEST"`
	TimeLimit     int    `env:"TIME_LIMIT"`
	BlockTime     int    `env:"BLOCK_TIME"`
	ServerPort    int    `env:"SERVER_PORT"`
}

func LogConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("rate_limiter")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
