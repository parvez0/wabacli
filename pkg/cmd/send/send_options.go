package send

import (
	"github.com/parvez0/wabacli/config"
	"github.com/spf13/cobra"
)

type Cmd string

type SendOptions struct {
	Config *config.Configuration
	File *Media
	VerifyContact bool
	Type string `validate: "required"`
	Message string
	Caption string
	UserId string `validate: "required"`
	MediaPath string
}

func NewSendOptions(c *config.Configuration) *SendOptions {
	return &SendOptions{
		Config: c,
	}
}

func (s *SendOptions) GetArgs() []string {
	var args []string
	for k, _ := range MediaTypeMapping {
		args = append(args, k)
	}
	return args
}

func (s *SendOptions) GetCommand(c string, cfg config.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use: c,

	}
	return cmd
}

func (s *SendOptions) Validate(args []string)  {

}
