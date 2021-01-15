package context

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
	"github.com/parvez0/wabacli/pkg/utils/validator"
)

type AddOptions struct {
	Config *config.Configuration
	Cluster config.Cluster
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
	derrs := validator.Validate(ap.Cluster)
	errs := validator.Validate(ap)
	errs = append(errs, derrs...)
	if len(errs) > 0 {
		handler.FatalError(fmt.Errorf("validating cluster parameters - %v", errs))
	}
}

func (ap *AddOptions) ResetPassword()  {
	if ap.Reset {
		log.Debug("validating fields before processing")
		if ap.NewPassword == "" {
			log.Error(badrequest.BadRequest{
				Code:        400,
				Title:       "Required field missing",
				Description: "NewPassword is required, if reset is enable",
			})
		}
	}
}

func (ap *AddOptions) Login()  {
	helpers.Login(ap)
}