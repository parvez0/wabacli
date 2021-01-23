package _switch

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	setLong = templates.LongDesc(i18n.T(`
		Switch the current context
		
		You can switch the current active context to manage multiple accounts at the same time.
	`))
	setExample = templates.Examples(i18n.T(`
        # Set current context
        wabactl context switch <number>
	`))
)

func NewDefaultSetCmd(c *config.Configuration) *cobra.Command {
	s := &SwitchOptions{
		Config: c,
	}
	cmd := &cobra.Command{
		Use: "switch",
		Long: setLong,
		Short: i18n.T("Switch the current context"),
		Run: setAccount(s),
		SuggestFor: []string{"set", "list", "switch", "activate"},
	}
	return cmd
}

func setAccount(s *SwitchOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		s.Validate(args)
		s.Run()
	}
}