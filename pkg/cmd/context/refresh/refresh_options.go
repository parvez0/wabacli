package refresh

import (
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
)

type RefOptions struct {
	Config *config.Configuration
	Password string
}

func NewRefreshOptions(c *config.Configuration) *RefOptions {
	return &RefOptions{
		Config:   c,
	}
}

// TODO: validate the request for refresh of multiple accounts
func (r *RefOptions) Validate() error {
	if r.Password == "" {
		return fmt.Errorf("password not provided")
	}
	return nil
}

func (r *RefOptions) Run()  {
	log.Debug("generating new auth token")
	auth := helpers.Login(&r.Config.CurrentCluster, r.Password, "")
	r.Config.CurrentCluster.Auth = auth
	err := r.Config.AddCluster(r.Config.CurrentCluster)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to refresh account: %v", err))
	}
}
