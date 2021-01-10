package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
	config *Configuration
)

// New provides a singleton for creating the configuration
// Once handles the cases where multiple routines are trying
// to initialize the config file
func New() (c *Configuration, err error) {
	if config != nil {
		return config, nil
	}
	once.Do(func() {
		c, err = initializeConfig()
	})
	return
}

// initializeConfig will read the config file from home directory
// and decodes it into Configuration structure
func initializeConfig() (*Configuration, error)  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.waba/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("clusters", []Cluster{
		{
			Auth: "",
			Server: "https://localhost",
			Name: "Default Server",
			Number: "",
			Context: "default",
			Insecure: true,
		},
	})
	viper.SetDefault("current_context", "default")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfig()
		} else {
			return nil, err
		}
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	for _, v := range config.Clusters {
		if v.Context == config.CurrentContext {
			config.CurrentCluster = v
			break
		}
	}
	return config, nil
}