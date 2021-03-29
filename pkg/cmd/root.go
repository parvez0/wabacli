package cmd

import (
	config2 "github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/cmd/context"
	"github.com/parvez0/wabacli/pkg/cmd/send"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"os"
)

var config, _ = config2.GetConfig()

// NewDefaultWabaCmdWithConfig returns root command for wabacli cli
func NewDefaultWabaCmdWithConfig() *cobra.Command {
	var v bool
	cmd := &cobra.Command{
		Use:   "wabacli",
		Short: "wabacli provides a cli to interact with whatsapp business account",
		Long:  templates.LongDesc(`
			wabacli provides a cli to interact with whatsapp business account

			you can find more information of the api's at:
            	https://developers.facebook.com/docs/whatsapp/api/account
		`),
		Run: func(cmd *cobra.Command, args []string) {
			if v {
				log.Info(config.Version)
				os.Exit(0)
			}
			cmd.Help()
		},
	}
	cmd.Flags().BoolVarP(&v, "version","v", false, "current version of wabacli")
	cmd.AddCommand(context.NewContextCommand(config))
	cmd.AddCommand(send.NewDefaultSendCmd(config))
	return cmd
}

func Execute()  {
	NewDefaultWabaCmdWithConfig().Execute()
}