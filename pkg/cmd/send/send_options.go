package send

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	handler2 "github.com/parvez0/wabacli/pkg/internal/handler"
	"github.com/parvez0/wabacli/pkg/utils/helpers"
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"github.com/parvez0/wabacli/pkg/utils/validator"
	"github.com/spf13/cobra"
	"os"
)

type Cmd string

// SendOptions groups all the basic types and
// data which are required to process the user
// request and send the message to user
type SendOptions struct {
	Config *config.Configuration
	File *types.Media
	Message types.WAMessage
	Request map[string]interface{}
	FilePath string
	Url string
	Json string
	VerifyContact bool
	VerifyAsync bool
	VerifyForced bool
}

// NewSendOptions initializes the SendOptions and returns a pointer to it
func NewSendOptions(c *config.Configuration) *SendOptions {
	return &SendOptions{
		Config: c,
	}
}

// GetCmdList returns an array of all the available commands
// which are currently supported by this cli
func (s *SendOptions) GetCmdList() []string {
	return []string{ "text", "image", "document", "video", "audio" }
}

// GetCommand generates a cobra command which will be added as a
// child command to the send command, it also defines all the ,
// required parameters that all the subcommands supports
func (s *SendOptions) GetCommand(arg string) *cobra.Command {
	cmd := &cobra.Command{
		Use: arg,
		Short: ShortDesc[arg],
		Long: LongDesc[arg],
		Example: ExampleDesc[arg],
		Run: s.getRunnable(arg),
	}
	cmd.Flags().BoolVarP(&s.VerifyContact, "verify", "v", true, "verify receivers contact before sending the message")
	cmd.Flags().BoolVarP(&s.VerifyAsync, "verify-async", "b", false, "don't wait for the verify contact response")
	cmd.Flags().BoolVarP(&s.VerifyForced, "verify-forced", "f", false, "forced verify the account")
	cmd.Flags().StringVarP(&s.Message.RecipientType, "recipient-type", "r", "individual", "recipient-type it can either be individual or group")
	cmd.Flags().IntVarP(&s.Message.To, "to", "t", 0, "receivers registered mobile number with country code")
	switch cmd.Use {
	case "text":
		cmd.Flags().StringVarP(&s.Message.Text.Body, "message", "m", "", "text message which needs to be send to receiver")
	default:
		cmd.Flags().StringVarP(&s.FilePath, "path", "p", "", "relative path to the file which will be send")
		cmd.Flags().StringVarP(&s.Url, "url", "u", "", "file url which needs to be sent to the user")
		cmd.Flags().StringVarP(&s.Message.Caption, "caption", "c", "", "caption to be added to the media file")
		cmd.Flags().BoolVarP(&s.Message.PreviewURL, "preview-url", "s", false, "preview url for showing the preview of the link inside a message")
	}
	return cmd
}

// Validate will verify if any required field is empty,
// if the required filed is not provided it will exit.
func (s *SendOptions) Validate()  {
	if s.Message.To == 0 {
		num, err := templates.NewPromptNumber()
		if err != nil {
			handler.FatalError(fmt.Errorf("failed to read mobile number: %s", err.Error()))
		}
		s.Message.To = num
	}
	err := validator.Validate(s)
	msgErr := validator.Validate(s.Message)
	err = append(err, msgErr ...)
	if len(err) > 0{
		handler.FatalError(handler.FormatError("validation failed", err))
	}
	if s.VerifyContact {
		log.Debug("verifying the contact before sending the message")
		resp, err := helpers.VerifyContact(&s.Config.CurrentCluster, s.Message.To, s.VerifyAsync, s.VerifyForced)
		if err != nil {
			handler.FatalError(err)
		}
		status, err := resp.GetStatus()
		if err != nil  {
			handler.FatalError(err)
		}
		if status == "invalid" {
			handler2.JsonResponse(resp)
			os.Exit(1)
		}
		log.Debug("contact verified successfully - ", resp)
	}
	log.Debug("field validation is successful, proceeding for command parsing")
}

// Parse further verifies the request body and initiates
// all the required functions like file reader which will
// return a Media object which provides the functionality
// to read and store binary data of a file
func (s *SendOptions) Parse()  {
	log.Debug("parsing the fields, to format request payload")
	switch s.Message.Type {
	case "text":
		if s.Message.Text.Body == "" {
			handler.FatalError(fmt.Errorf("validation failed: ValidationError(RequiredFiled) missing required field \"Message;\""))
		}
	default:
		if s.Url != "" {
			return
		}
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

// Run executes the commands by uploading the file
// in case of media message and then sending the
// message to the user, if any error it will print
func (s *SendOptions) Run() {
	mediaId := ""
	if s.File != nil {
		mediaId = s.uploadMedia()
	}
	s.processBody(mediaId)
	resp := s.sendMessage()
	handler2.JsonResponse(resp)
}

// getRunnable will return the func which can be used to run the specified command
// this function is required to avoid s.Message.Type being overwritten by other sub
// commands
func (s *SendOptions) getRunnable(arg string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		s.Message.Type = arg
		s.Validate()
		s.Parse()
		s.Run()
	}
}

// uploadMedia calls the helper UploadMedia function to upload
// the file and generate the mediaId if any error it will exit
// with the error received by the api response
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

// processBody is an internal function which will remove all the
// not used fields from the Message struct then it will add the
// to SendOptions as Request field which will be used for sending
// the message through api
func (s *SendOptions) processBody(mediaId string) {
	if s.Json == "" {
		switch s.Message.Type {
		case "audio":
			s.Message.Audio.ID = mediaId
			s.Message.Audio.Caption = s.Message.Caption
			s.Message.Audio.Link = s.Url
		case "video":
			s.Message.Video.ID = mediaId
			s.Message.Video.Caption = s.Message.Caption
			s.Message.Video.Link = s.Url
		case "document":
			s.Message.Document.ID = mediaId
			s.Message.Document.Caption = s.Message.Caption
			s.Message.Document.Link = s.Url
		case "image":
			s.Message.Image.ID = mediaId
			s.Message.Image.Caption = s.Message.Caption
			s.Message.Image.Link = s.Url
		}
	}
	// Marshall the message object to get string
	m, _ := json.Marshal(s.Message)
	var finalBody map[string]interface{}

	json.Unmarshal(m, &finalBody)
	for _, v := range s.GetCmdList(){
		if s.Message.Type == v {
			continue
		}
		delete(finalBody, v)
	}
	log.Debug(fmt.Sprintf("after formatting the request body - %+v", finalBody))
	s.Request = finalBody
}

// sendMessage calls the helper function SendMessage
// for sending the message to the user, it will use
// the Request field which was processed in the previous
// step as request body and it will exit if encountered
// any error by print the json output to stdin
func (s *SendOptions) sendMessage() string {
	byts, err := helpers.SendMessage(&s.Config.CurrentCluster, &s.Request)
	if err != nil {
		handler.FatalError(err)
	}
	log.Debug("message send triggered response received - ", string(byts))
	return string(byts)
}
