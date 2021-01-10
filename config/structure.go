package config

// Configuration defines the basic structure of config file
// Config can be provided in json, yaml or from environment
type Configuration struct {
	Clusters []Cluster `json:"clusters",yaml:"clusters"`
	CurrentContext string `json:"current_context",yaml:"current_context"`
	CurrentCluster Cluster `json:"current_cluster",yaml:"current_cluster"`
}

// Cluster holds the basic information of the connected cluster
// Auth is authentication generated after login
// Server is the base url of the account
// Name (optional) is the name of the cluster
// Number is the phone number of the cluster
// VerifySSL will define if ssl needs to be verified during the
// api call, defaults to true
type Cluster struct {
	Auth string
	Server string
	Name string
	Number string
	Insecure bool
	Context string
}
