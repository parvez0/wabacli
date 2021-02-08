package send

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	sendLong = templates.LongDesc(i18n.T(`
	Send message to user
	
	You can send text, media videos to the user. Just provide the user id and message to send
	`))
	sendExample = templates.Examples(i18n.T(`
	# send text message
	wabactl send text --message "Hello World"
	 `))
)

// NewDefaultSendCmd returns a cobra command with all of it's
// child commands added to it, send command supports all the
// basic type of messages a user can send like text, images,
// video etc which are supported by whatsapp. also send command
// is also capable of taking json request body and directly
// processing the request without using it's children.
func NewDefaultSendCmd(c *config.Configuration) *cobra.Command {
	s := NewSendOptions(c)
	cmd := &cobra.Command{
		Use:                        "send",
		SuggestFor:                 []string{"text", "message"},
		Short:                      i18n.T("Send message to user"),
		Long:                       sendLong,
		Example:                    sendExample,
		Run:                        getDefaultSendCmd(s),
	}
	avaCmds := s.GetCmdList()
	for _, cd := range avaCmds {
		cmd.AddCommand(s.GetCommand(cd))
	}
	return cmd
}

func getDefaultSendCmd(s *SendOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	}
}