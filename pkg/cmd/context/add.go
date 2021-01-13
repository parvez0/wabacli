package context

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	addLong = templates.LongDesc(i18n.T(`
		Add a new cluster
		
		You can add multiple cluster to config and keep track of all of your servers at a single place.
	`))
	addExample = templates.Examples(i18n.T(`
        # Add a new cluster with parameters
        wabactl context add --server <IP> --cluster-name <string> --number <without country code> --country-code <91> --username <string> --password 

        # List info about a single cluster
        wabactl context add --json "{cluster_name: '', number: '', country_code: '', username: '', password: ''}"
	`))
)

func NewDefaultAddCmd(c *config.Configuration) *cobra.Command {
	ap := NewAddOptions(c)
	cmd := &cobra.Command{
		Use: "add [--cluster-name|-cn, --number|-n, --country_code|-cc, --username|-u, --password|-p]",
		Short: i18n.T("Add a new cluster"),
		Long: addLong,
		Example: addExample,
		Run: addAccount(ap),
	}
	cmd.Flags().StringVarP(&ap.Server, "server", "s", "https://localhost", "whatsapp infra server address")
	cmd.Flags().StringVarP(&ap.Json, "json", "js", "", "json object string with all information")
	cmd.Flags().StringVarP(&ap.ClusterName, "cluster-name", "cn", "", "name for your cluster entry in config file")
	cmd.Flags().StringVarP(&ap.Number, "number", "n", "", "whatsapp account number connected to cluster without country code")
	cmd.Flags().StringVarP(&ap.CountryCode, "country-code", "cc", "", "assigned country code")
	cmd.Flags().StringVarP(&ap.Username, "username", "u", "admin", "whatsapp account admin username")
	cmd.Flags().StringVarP(&ap.Password, "password", "p", "", "whatsapp account admin password")
	cmd.Flags().BoolVarP(&ap.Reset, "reset", "r", false, "reset initial password, if specified new_password is required")
	cmd.Flags().StringVarP(&ap.NewPassword, "new-password", "np", "", "whatsapp account new admin password after reset")
	return cmd
}

func addAccount(ap *AddOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		ap.Parse()
		ap.Validate()
		ap.ResetPassword()
		ap.Login()
	}
}