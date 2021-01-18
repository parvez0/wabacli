package config

import (
	"fmt"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"regexp"
)

// Configuration defines the basic structure of config file
// Config can be provided in json, yaml or from environment
type Configuration struct {
	Clusters []Cluster `mapstructure:"clusters" json:"clusters"`
	CurrentContext string `mapstructure:"current_context" json:"current_context"`
	CurrentCluster Cluster `mapstructure:"current_cluster" json:"current_cluster"`
}

// Cluster holds the basic information of the connected cluster
// Auth is authentication generated after login
// Server is the base url of the account
// Name (optional) is the name of the cluster
// CountryCode is the officially assigned code
// Number is the phone number of the cluster
// VerifySSL will define if ssl needs to be verified during the
// api call, defaults to true
type Cluster struct {
	Auth string
	Name string `json:"cluster_name"`
	Context string
	CountryCode int `json:"country_code,int" validate:"required"`
	Number int `json:"number,int" validate:"required"`
	Server string `json:"server" validate:"required"`
	Username string `json:"username" validate:"required"`
	Insecure bool
}

func (c *Configuration) MatchAcc(acc string) (int, []string) {
	// creating a regexp for the query provided
	reg, err := regexp.Compile(acc)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to create regex exp: %v", err))
	}
	var accs []string
	var l = 0
	// getting all the matched accounts from the cluster
	for _, v := range c.Clusters {
		if matched := reg.MatchString(v.Context); matched {
			accs = append(accs, v.Context)
			l++
		}
	}
	return l, accs
}

// UpdateConfig will update the config file to latest
// after addition or removal of an account from the cluster
func (c *Configuration) AddCluster(cluster Cluster) error {
	// removing the default element from the slice
	c.Clusters = removeElement(c.Clusters, DefaultCurrentContext, 0)
	c.Clusters = removeElement(c.Clusters, "", cluster.Number)
	c.Clusters = append(c.Clusters, cluster)
	vp.Set("current_context", cluster.Context)
	vp.Set("current_cluster", cluster)
	vp.Set("clusters", c.Clusters)
	_ = vp.WriteConfig()
	return nil
}

func (c *Configuration) RemoveCluster(context string) error {
	// removing the default element from the slice
	c.Clusters = removeElement(c.Clusters, context, 0)
	if len(c.Clusters) == 0 {
		return nil
	}
	vp.Set("current_context", c.Clusters[0].Context)
	vp.Set("current_cluster", c.Clusters[0])
	vp.Set("clusters", c.Clusters)
	_ = vp.WriteConfig()
	return nil
}

func (c *Configuration) UpdateContext(context string) error {
	for _, v := range c.Clusters {
		if v.Context == context {
			vp.Set("current_context", v.Context)
			vp.Set("current_cluster", v)
			_ = vp.WriteConfig()
			break
		}
	}
	return nil
}

func removeElement(c []Cluster, ele string, num int) (clus []Cluster) {
	for i, v := range c {
		if v.Context == ele || v.Number == num {
			c[len(c) - 1], c[i] = c[i], c[len(c) - 1]
			clus = c[:len(c) - 1]
			return
		}
	}
	clus = c
	return
}