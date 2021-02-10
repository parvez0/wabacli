package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/spf13/viper"
	"os"
	"sync"
)

const (
	DefaultCurrentContext = "default"
	DefaultServer = "https://localhost"
)

var (
	once sync.Once
	config *Configuration
	vp *viper.Viper

)

// TODO add auth expiry time to config
// New provides a singleton for creating the configuration
// Once handles the cases where multiple routines are trying
// to initialize the config file
func GetConfig() (c *Configuration, err error) {
	if config != nil {
		return config, nil
	}
	once.Do(func() {
		c, err = initializeConfig()
	})
	return
}

// createConfig will helps in creating the config
// file if it doesn't exits. it initializes a empty
// config file which can used to write the data of
// the accounts which will be added
func createConfig()  {
	cp := "/etc/wabactl"
	home, err := os.UserHomeDir()
	if err != nil {
		log.Debug(fmt.Sprintf("home directory not found %s. using '%s/config' directory as default config path", err.Error(), cp))
	} else {
		cp = home + "/.waba"
	}
	err = os.Mkdir(cp, 0700)
	if ex := os.IsNotExist(err); ex {
		handler.FatalError(fmt.Errorf("failed to create config file at '%s'; %s", cp, err.Error()))
	}
	_, err = os.Stat(cp + "/config")
	if err != nil {
		log.Debug("creating config file at :", cp)
		_, err = os.Create(cp + "/config")
		if err != nil {
			handler.FatalError(fmt.Errorf("failed to create config file at '%s'; %s", cp, err.Error()))
		}
	}
}

// initializeConfig will read the config file from home directory
// and decodes it into Configuration structure
func initializeConfig() (*Configuration, error)  {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("$HOME/.waba/")
	vp.AddConfigPath("/etc/wabacli")
	vp.AddConfigPath(".")
	vp.AutomaticEnv()

	vp.SetDefault("clusters", []Cluster{
		{
			Auth: "",
			Server: DefaultServer,
			Name: "Default Server",
			Number: 987654321,
			Context: DefaultCurrentContext,
			Insecure: true,
		},
	})

	vp.SetDefault("current_context", DefaultCurrentContext)
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfig()
		} else {
			return nil, err
		}
	}
	err := vp.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	if config.CurrentContext != config.CurrentCluster.Context {
		for _, v := range config.Clusters {
			if v.Context == config.CurrentContext {
				config.CurrentCluster = v
				break
			}
		}
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		log.Debug("config change detected: ", in.Name, in.Op.String(), ", updating to latest")
		_ = vp.WriteConfig()
	})
	return config, nil
}