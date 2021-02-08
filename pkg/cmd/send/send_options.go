package send

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"github.com/parvez0/wabacli/pkg/utils/validator"
	"github.com/spf13/cobra"
)

type Cmd string

type SendOptions struct {
	Config *config.Configuration
	File *types.Media
	Message types.WAMessage
	VerifyContact bool
	FilePath string
}

func NewSendOptions(c *config.Configuration) *SendOptions {
	return &SendOptions{
		Config: c,
	}
}

func (s *SendOptions) GetCmdList() []string {
	return []string{ "text", "image", "document", "video", "audio" }
}

func (s *SendOptions) GetCommand(arg string, cfg *config.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use: arg,
		Short: ShortDesc[arg],
		Long: LongDesc[arg],
		Example: ExampleDesc[arg],
		Run: s.getRunnable(arg),
	}
	cmd.Flags().BoolVarP(&s.VerifyContact, "verify", "v", true, "verify receivers contact before sending the message")
	cmd.Flags().StringVarP(&s.Message.RecipientType, "recipient-type", "r", "individual", "recipient-type it can either be individual or group")
	cmd.Flags().StringVarP(&s.Message.To, "to", "t", "", "receivers registered mobile number with country code")
	switch cmd.Use {
	case "text":
		cmd.Flags().StringVarP(&s.Message.Text.Body, "message", "m", "", "text message which needs to be send to receiver")
	default:
		cmd.Flags().StringVarP(&s.FilePath, "path", "p", "", "relative path to the file which will be send")
		cmd.Flags().StringVarP(&s.Message.Caption, "caption", "c", "", "caption to be added to the media file")
		cmd.Flags().BoolVarP(&s.Message.PreviewURL, "preview-url", "s", true, "preview url for video files, by default it is enable")
	}
	return cmd
}

func (s *SendOptions) Validate()  {
	err := validator.Validate(s)
	msgErr := validator.Validate(s.Message)
	err = append(err, msgErr ...)
	if len(err) > 0{
		handler.FatalError(handler.FormatError("validation failed", err))
	}
	log.Debug("field validation is successful, proceeding for command parsing")
}

func (s *SendOptions) Parse()  {
	log.Debug("parsing the fields, to format request payload")
	switch s.Message.Type {
	case "text":
		if s.Message.Text.Body == "" {
			handler.FatalError(fmt.Errorf("validation failed: ValidationError(RequiredFiled) missing required field \"Message;\""))
		}
	default:
		if s.FilePath == "" {
			handler.FatalError(fmt.Errorf("validation failed: ValidationError(RequiredFiled) missing required field \"Path;\""))
		}
		var err error
		// initializing a new file reader object
		s.File, err = types.NewFileReader(s.FilePath)
		if err != nil {
			handler.FatalError(err)
		}
		err = s.File.Read()
		if err != nil {
			handler.FatalError(err)
		}
	}
}

func (s *SendOptions) Run() {

}

func (s *SendOptions) uploadMedia() string {
	log.Debug("uploading file to whatsapp for generating media id")
	byts, err := helpers.UploadMedia(&s.Config.CurrentCluster, s.File)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to upload file: %v", err))
	}
	var resp types.MediaResponse
	err = json.Unmarshal(byts, &resp)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to parse media response: %v", err))
	}
	mediaId, err := resp.GetId()
	if err != nil {
		handler.FatalError(err)
	}
	log.Debug("file uploaded successfully, generated media id :", mediaId)
	return mediaId
}

func (s *SendOptions) sendMessage()  {

}

func (s *SendOptions) getRunnable(arg string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		s.Message.Type = arg
		s.Validate()
		s.Parse()
		s.Run()
	}
}
