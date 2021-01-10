package cmd

import (
	config2 "github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
)

var config, _ = config2.GetConfig()

func newDefaultWabaCmdWithConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wabactl",
		Short: "wabactl provides a cli to interact with whatsapp business account",
		Long:  templates.LongDesc(`
			wabactl provides a cli to interact with whatsapp business account

			you can find more information of the api's at:
            	https://developers.facebook.com/docs/whatsapp/api/account
`),
		Run:   nil,
	}
	return cmd
}

func Execute()  {
	newDefaultWabaCmdWithConfig().Execute()
}