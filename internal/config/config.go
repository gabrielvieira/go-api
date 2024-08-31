package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	DBUrl                          string        `mapstructure:"db_url"`
	DBUser                         string        `mapstructure:"db_user"`
	DBPassword                     string        `mapstructure:"db_password"`
	DBSchema                       string        `mapstructure:"db_schema"`
	DBMaxConnections               int           `mapstructure:"db_max_connections"`
	DBMaxIdleConnections           int           `mapstructure:"db_max_idle_connections"`
	DBConnectionMaxLifetimeSeconds time.Duration `mapstructure:"db_connection_max_lifetime_seconds"` // int64

	APIPort string `mapstructure:"api_port"`
}

func New() (*Config, error) {

	var config Config

	viper.SetConfigName(getConfigName())
	viper.AddConfigPath("configs")
	viper.SetDefault("api_port", 9090)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigName() string {
	appEnv := os.Getenv("APPLICATION_ENVIRONMENT")
	if appEnv == "" {
		appEnv = "dev"
	}
	return appEnv
}
