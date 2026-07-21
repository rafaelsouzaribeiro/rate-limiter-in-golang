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

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
