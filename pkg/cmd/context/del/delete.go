package del

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	delLong = templates.LongDesc(i18n.T(`
		Delete a cluster from config
		
		You can remove an account from the present config file.
	`))

	delExample = templates.Examples(i18n.T(`
        # Set current context
        wabacli context del <number | name>
	`))
)

func NewDefaultDeleteCmd(c *config.Configuration) *cobra.Command {
	d := NewDeleteOptions(c)
	cmd := &cobra.Command{
		Use: "delete",
		Long: delLong,
		Example: delExample,
		Short: i18n.T("Switch the current context"),
		Run: DeleteAccount(d),
		SuggestFor: []string{"remove"},
	}
	return cmd
}

func DeleteAccount(d *DelOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		d.Validate(args[0])
		d.Run()
	}
}
