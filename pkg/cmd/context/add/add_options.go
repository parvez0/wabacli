package add

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
	"github.com/parvez0/wabacli/pkg/utils/validator"
	"strings"
)

type AddOptions struct {
	Config *config.Configuration
	Cluster config.Cluster
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
		handler.FatalError(fmt.Errorf("json parse failed: %v", err))
	}
	err = json.Unmarshal([]byte(ap.Json), &ap.Cluster)
	if err != nil {
		handler.FatalError(fmt.Errorf("json parse failed: %v", err))
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
	slug := strings.ToLower(strings.Join(strings.Split(ap.Cluster.Name, " "), "_"))
	ap.Cluster.Context = fmt.Sprintf("%d_%s", ap.Cluster.Number, slug)
}

func (ap *AddOptions) ResetPassword()  {
	if ap.Reset {
		log.Debug("validating fields before processing")
		if ap.NewPassword == "" {
			handler.FatalError(&badrequest.BadRequest{
				Code:        400,
				Title:       "Required field missing",
				Description: "NewPassword is required, if reset is enable",
			})
		}
	}
}

func (ap *AddOptions) Login()  {
	ap.Cluster.Auth = helpers.Login(&ap.Cluster, ap.Password, ap.NewPassword)
	ap.save()
}

func (ap *AddOptions) save()  {
	ap.Config.AddCluster(ap.Cluster)
}