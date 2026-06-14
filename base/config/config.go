package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Google struct {
		ClientID     string `mapstructure:"client_id" json:"client_id"`
		ClientSecret string `mapstructure:"client_secret" json:"client_secret"`
	} `mapstructure:"google" json:"google"`
	Database struct {
		Uri                      string `mapstructure:"uri" json:"uri"`
		MaxOpenConns             int    `mapstructure:"max_open_conns" json:"max_open_conns"`
		MaxIdleConns             int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
		ConnMaxLifetimeMins      int    `mapstructure:"conn_max_lifetime_mins" json:"conn_max_lifetime_mins"`
		ConnectionTimeoutSeconds int    `mapstructure:"connection_timeout_seconds" json:"connection_timeout_seconds"`
	} `mapstructure:"database" json:"database"`
	JWT struct {
		Secret          string `mapstructure:"secret" json:"secret"`
		ExpirationHours int    `mapstructure:"expiration_hours" json:"expiration_hours"`
	} `mapstructure:"jwt" json:"jwt"`
}

type config Configuration

var Config *Configuration

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	var tempConfig Configuration
	if err := viper.Unmarshal(&tempConfig); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	Config = &tempConfig
	return nil
}
