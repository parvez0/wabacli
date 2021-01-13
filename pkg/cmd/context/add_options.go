package context

import (
	"encoding/json"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"reflect"
)

type AddOptions struct {
	Config *config.Configuration
	Server string
	ClusterName string
	CountryCode string
	Number string
	Username string
	Password string
	Reset bool
	NewPassword string
	Json string
}

func NewAddOptions(c *config.Configuration) *AddOptions {
	return &AddOptions{
		Config: c,
	}
}

func (ap *AddOptions) Parse()  {
	if ap.Json == "" {
		return
	}
	log.Debug("parsing json object into struct")
	err := json.Unmarshal([]byte(ap.Json), ap)
	if err != nil {
		log.Error("json parse failed: ", err)
	}
	log.Debug("parsed values into struct fields: ", ap)
}


func (ap *AddOptions) Validate()  {
	log.Debug("validating fields before processing")
	val := reflect.ValueOf(ap)
	for i := 0; i < val.NumField(); i++ {

	}
}