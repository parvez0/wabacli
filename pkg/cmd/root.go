package cmd

import (
	config2 "github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/cmd/context"
	"github.com/parvez0/wabacli/pkg/cmd/send"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
)

var config, _ = config2.GetConfig()


// newDefaultWabaCmd
func NewDefaultWabaCmdWithConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wabactl",
		Short: "wabactl provides a cli to interact with whatsapp business account",
		Long:  templates.LongDesc(`
			wabactl provides a cli to interact with whatsapp business account

			you can find more information of the api's at:
            	https://developers.facebook.com/docs/whatsapp/api/account
		`),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(context.NewContextCommand(config))
	cmd.AddCommand(send.NewDefaultSendCmd(config))
	return cmd
}

func Execute()  {
	NewDefaultWabaCmdWithConfig().Execute()
}