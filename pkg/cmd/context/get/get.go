package get

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	getLong = templates.LongDesc(i18n.T(`
		Display cluster info
		
		Prints the table information about the clusters added to the context.
		You can also apply some selectors by specifying unique filed name like
        --selector available fields for selector are name, number and context.
	`))
	getExample = templates.Examples(i18n.T(`
        # List info of all clusters
        wabactl context get

        # List info about a single cluster
        wabactl context get --name example
	`))
)

func NewGetOptions(c *config.Configuration) *GetOptions {
	return &GetOptions{
		Config:         c,
	}
}

func NewGetCmdWithConfig(c *config.Configuration) *cobra.Command {
	opts := NewGetOptions(c)
	cmd := &cobra.Command{
		Use:        "get [(-s|--selector=<name, number>)] [flags]",
		Long:       getLong,
		Short: 		i18n.T("Display cluster info"),
		Example:    getExample,
		Run:        getAccountsWithOptions(opts),
		SuggestFor: []string{"list", "accounts"},
	}
	cmd.Flags().StringVarP(&opts.Selector,"selector", "s", "", i18n.T("unique value which can be used to fetch account, it can be name or number without country code"))
	return cmd
}

func getAccountsWithOptions(o *GetOptions) func(command *cobra.Command, args []string) {
	return func (cmd *cobra.Command, args []string) {
		getAccountsWithFlags(o)
	}
}

func getAccountsWithFlags(o *GetOptions) {
	o.Validate()
	o.Run()
}