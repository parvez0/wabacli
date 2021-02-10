package send

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	sendLong = templates.LongDesc(i18n.T(`
	Send message to user
	
	You can send text, media videos to the user. Just provide the user id and message to send. If you want
	more control over the request just create a json request body and send it directly to whatsapp
	`))
	sendExample = templates.Examples(i18n.T(`
	# send text message
	wabacli send text --message "Hello World"
	
	# send message using json
	wabacli send --json '{"to":"{{Recipient-WA-ID}}","type":"text","recipient_type":"individual","text":{"body":"<Message Text>"}}'
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
	cmd.Flags().StringVarP(&s.Json, "json", "j", "", "request body of the whole request with json")
	return cmd
}

func getDefaultSendCmd(s *SendOptions) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if s.Json == "" {
			cmd.Help()
			return
		}
		err := json.Unmarshal([]byte(s.Json), &s.Message)
		if err != nil {
			handler.FatalError(fmt.Errorf("failed to parse request body - %s", err.Error()))
		}
		s.Run()
	}
}