package context

import (
	"encoding/json"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"github.com/parvez0/wabacli/pkg/utils/validator"
	"reflect"
)

type AddOptions struct {
	Config *config.Configuration
	Cluster *config.Cluster
	Password string `validate:"required"`
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
	errs := validator.Validate(ap)
	if len(errs) > 0 {
		log.Error(errs)
	}
}

func (ap *AddOptions) ResetPassword()  {
	if ap.Reset {
		log.Debug("validating fields before processing")
		if ap.NewPassword == "" {
			log.Error(badrequest.BadRequest{
				Code:        400,
				Title:       "Required field missing",
				Description: "new_password is required, if reset is enable",
			})
		}

	}
}

func (ap *AddOptions) Login()  {
	log.Debug("validating fields before processing")
	val := reflect.ValueOf(ap)
	for i := 0; i < val.NumField(); i++ {

	}
}