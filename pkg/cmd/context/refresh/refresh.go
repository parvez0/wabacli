package refresh

import (
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	refreshLong = templates.LongDesc(i18n.T(`
		Refresh auth token
		
		Refresh the auth token by providing password, you can provide multiple accounts to be refreshed at same time.
		By default it will refresh the current context
	`))
	refreshExample = templates.Examples(i18n.T(`
        # List info of all clusters
        wabactl context refresh

        # List info about a single cluster
        wabactl context refresh 
	`))
)

func NewRefreshCmd(c *config.Configuration) *cobra.Command {
	opts := NewRefreshOptions(c)
	cmd := &cobra.Command{
		Use:        "refresh",
		Long:       refreshLong,
		Short: 		i18n.T("Refresh current cluster"),
		Example:    refreshExample,
		Run:        refreshAccountsWithOptions(opts),
		SuggestFor: []string{"list", "accounts"},
	}
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "admin account password for login into the whatsapp infra")
	return cmd
}

func refreshAccountsWithOptions(o *RefOptions) func(command *cobra.Command, args []string) {
	return func (cmd *cobra.Command, args []string) {
		err := o.Validate()
		if err != nil {
			res, err := templates.NewPromptPassword()
			if err != nil {
				handler.FatalError(fmt.Errorf("failed to refresh account: %v", err))
			}
			o.Password = res
		}
		o.Run()
	}
}
