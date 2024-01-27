// File: config/flags.go

package config

import (
	"os"

	"github.com/spf13/viper"
)

// Config holds the configuration values for the application
type Config struct {
	ConfigFile string
}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gcpcisaudit-cli")
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err == nil {
		// If a config file is found, read it in.
		if err := viper.Unmarshal(&config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
