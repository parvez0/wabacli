package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/parvez0/wabacli/log"
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

func createConfigDirectory()  {
	cp := "/etc/wabactl"
	home, err := os.UserHomeDir()
	if err != nil {
		log.Debug(fmt.Sprintf("home directory not found %s. using '%s/config.yml' directory as default config path", err.Error(), cp))
	} else {
		cp = home + "/.waba"
	}
	err = os.Mkdir(cp, 0700)
	if ex := os.IsNotExist(err); ex {
		log.Error(fmt.Sprintf("failed to create config file at '%s'; %s", cp, err.Error()))
	}
}

// initializeConfig will read the config file from home directory
// and decodes it into Configuration structure
func initializeConfig() (*Configuration, error)  {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yml")
	vp.AddConfigPath("$HOME/.waba/")
	vp.AddConfigPath("/etc/wabactl")
	vp.AddConfigPath(".")
	vp.AutomaticEnv()

	vp.SetDefault("clusters", []Cluster{
		{
			Auth: "",
			Server: DefaultServer,
			Name: "Default Server",
			Number: "",
			Context: DefaultCurrentContext,
			Insecure: true,
		},
	})

	vp.SetDefault("current_context", DefaultCurrentContext)
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfigDirectory()
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
		vp.WriteConfig()
	})
	return config, nil
}

func removeDefaultElement(c []Cluster) (clus []Cluster) {
	for i, v := range c {
		if v.Context == DefaultCurrentContext {
			c[len(c) - 1], c[i] = c[i], c[len(c) - 1]
			clus = c[:len(c) - 1]
			return
		}
	}
	clus = c
	return
}

func UpdateConfig(cluster Cluster) error {
	c, err := GetConfig()
	if err != nil {
		return err
	}
	clus := removeDefaultElement(c.Clusters)
	c.Clusters = append(clus, cluster)
	vp.Set("current_context", cluster.Context)
	vp.Set("current_cluster", cluster)
	vp.Set("clusters", c.Clusters)
	vp.WriteConfig()
	return nil
}