package context

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/cmd/context/add"
	"github.com/parvez0/wabacli/pkg/cmd/context/get"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	contextLong = templates.LongDesc(i18n.T(`
		Current context manipulation
		
		Allows you to manipulate cli configuration to use multiple accounts at the same time, 
		you can add new accounts, delete, update them and more.
	`))

	contextExample = templates.Examples(i18n.T(`
		# List all accounts in current configuration.
		wabactl context list

		# Get info about the current .
		wabactl context get --selector 
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
	return cmd
}
