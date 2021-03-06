package context

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/cmd/context/add"
	"github.com/parvez0/wabacli/pkg/cmd/context/change"
	delete2 "github.com/parvez0/wabacli/pkg/cmd/context/del"
	"github.com/parvez0/wabacli/pkg/cmd/context/get"
	"github.com/parvez0/wabacli/pkg/cmd/context/refresh"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	contextLong = templates.LongDesc(i18n.T(`
		Current context manipulation
		
		Allows you to manipulate cli configuration to use multiple accounts at the same time, 
		you can add new accounts, del, update them and more.
	`))

	contextExample = templates.Examples(i18n.T(`
		# List all accounts in current configuration.
		wabacli context list

		# Get info about the current .
		wabacli context get --selector 
	`))
)

// TODO : add a global flag for request timeout
func NewContextCommand(c *config.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use: "context",
		Short: "Current context manipulation",
		Long: contextLong,
		Example: contextExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			return
		},
	}
	cmd.AddCommand(get.NewGetCmdWithConfig(c))
	cmd.AddCommand(add.NewDefaultAddCmd(c))
	cmd.AddCommand(change.NewDefaultSetCmd(c))
	cmd.AddCommand(delete2.NewDefaultDeleteCmd(c))
	cmd.AddCommand(refresh.NewRefreshCmd(c))
	return cmd
}
