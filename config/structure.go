package config

// Configuration defines the basic structure of config file
// Config can be provided in json, yaml or from environment
type Configuration struct {
	Clusters []Cluster `json:"clusters" yaml:"clusters"`
	CurrentContext string `json:"current_context" yaml:"current_context"`
	CurrentCluster Cluster `json:"current_cluster" yaml:"current_cluster"`
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
	Name string
	Context string
	CountryCode int `json:"country_code,int" validate:"required"`
	Number int `json:"number,int" validate:"required"`
	Server string `json:"server" validate:"required"`
	Username string `json:"username" validate:"required"`
	Insecure bool
}
