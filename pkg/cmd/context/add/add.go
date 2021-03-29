package add

import (
	"github.com/parvez0/wabacli/config"
	log2 "github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"log"
	"os"
	"strconv"
)

var (
	addLong = templates.LongDesc(i18n.T(`
		Add a new cluster
		
		You can add multiple cluster to config and keep track of all of your servers at a single place.
	`))
	addExample = templates.Examples(i18n.T(`
        # Add a new cluster with parameters
        wabacli context add --server <IP> --cluster-name <string> --number <without country code> --country-code <91> --username <string> --password 

        # List info about a single cluster
        wabacli context add --json "{cluster_name: '', number: '', country_code: '', username: '', password: ''}"
	`))
)

func NewDefaultAddCmd(c *config.Configuration) *cobra.Command {
	ap := NewAddOptions(c)
	cmd := &cobra.Command{
		Use:     "add [--cluster-name|-g, --number|-n, --country-code|-c, --username|-u, --password|-p, --json|-j]",
		Short:   i18n.T("Add a new cluster"),
		Long:    addLong,
		Example: addExample,
		Run:     addAccount(ap),
	}
	cmd.Flags().StringVarP(&ap.Json, "json", "j", "", "json object string with all information")
	cmd.Flags().StringVarP(&ap.Cluster.Server, "server", "s", "https://localhost", "whatsapp infra server address")
	cmd.Flags().StringVarP(&ap.Cluster.Name, "cluster-name", "g", "", "name for your cluster entry in config file")
	cmd.Flags().IntVarP(&ap.Cluster.Number, "number", "n", 0, "whatsapp account number connected to cluster without country code")
	cmd.Flags().IntVarP(&ap.Cluster.CountryCode, "country-code", "c", 0, "assigned country code")
	cmd.Flags().StringVarP(&ap.Cluster.Username, "username", "u", "admin", "whatsapp account admin username")
	cmd.Flags().StringVarP(&ap.Password, "password", "p", "", "whatsapp account admin password")
	cmd.Flags().BoolVarP(&ap.Reset, "reset", "r", false, "reset initial password, if specified new_password is required")
	cmd.Flags().StringVarP(&ap.NewPassword, "new-password", "o", "", "whatsapp account new admin password after reset")
	cmd.Flags().BoolVarP(&ap.Cluster.Insecure, "insecure", "i", false, "keep it true if you are using self generated tls")
	return cmd
}

func addAccount(ap *AddOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		ap.Parse()
		ap.Validate()
		if ap.Cluster.Number == 9876543216 {
			if num := os.Getenv("INFRA_NUMBER"); num != "" {
				ap.Cluster.Number, _ = strconv.Atoi(num)
			}
		}
		if ap.Password == "" {
			pwd, err := templates.NewPromptPassword()
			if err != nil {
				log.Fatal("failed to read password please try again")
			}
			ap.Password = pwd
		} else {
			log2.Warn("Using a password on the command line interface can be insecure.")
		}
		ap.ResetPassword()
		ap.Login()
	}
}