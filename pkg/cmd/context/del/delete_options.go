package del

import (
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"os"
)

type DelOptions struct {
	Config *config.Configuration
	Account string
}

func NewDeleteOptions(c *config.Configuration) *DelOptions {
	return &DelOptions{Config: c}
}

func (d *DelOptions) Validate(acc string) {
	l, accs := d.Config.MatchAcc(acc)
	if l == 0 {
		handler.FatalError(fmt.Errorf("account '%s' doesn't exits", acc))
	}

	var res string
	var err error

	if l > 0 && accs[0] != acc {
		res, err = templates.NewPromptSelect("Delete Account", accs)
		if err != nil {
			handler.FatalError(fmt.Errorf("failed to del account : %v", err))
		}
	}

	ans, err := templates.NewPromptSelect(fmt.Sprintf("Do you really want to delete '%s'", res), []string{"Yes", "No"})
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to del account : %v", err))
	}
	if ans == "No" {
		os.Exit(0)
	}
	d.Account = res
}

func (d *DelOptions) Run()  {
	err := d.Config.RemoveCluster(d.Account)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to remove account: %v", err))
	}
}
