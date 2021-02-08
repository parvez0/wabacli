package send

import (
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	// ShortDesc defines a map for all the short description
	// of the available commands
	ShortDesc = map[string]string {
		"text": i18n.T("Send a text message"),
		"image": i18n.T("Send an image file"),
		"document": i18n.T("Send a document"),
		"audio": i18n.T("Send an audio file"),
		"video": i18n.T("Send a video file"),
	}
	// LongDesc defines a map for all the Long description
	// of the available commands which explain the function
	// of each of the subcommand
	LongDesc = map[string]string{
		"text": templates.LongDesc(i18n.T(`
			Send a text message
			
			send allows you to send a text message to a user,  it sends the message
			to the whatsapp infrastructure from there a job will be created which
			will take care of sending the message to the user
		`)),
		"image": templates.LongDesc(i18n.T(`
			Send an image file
			
			send a jpeg or a png file to the user, you need to provide the path to the file
			or you can directly give the url, from there it will read and uploaded it to the
			whatsapp infra, after that it will send to the user
		`)),
		"document": templates.LongDesc(i18n.T(`
			Send an document
			
			send a document supported by whatsapp, you need to provide the path to the file 
			or you can directly give the url, from there it will read and uploaded to the
			whatsapp infra, after that it will send to the user
		`)),
		"audio": templates.LongDesc(i18n.T(`
			Send an audio file
			
			send an audio file (.mp3 or .mp4) supported by whatsapp, you need to provide the path to the file
			from there it will read and uploaded it to the whatsapp infra, after that it will 
			send to the user
		`)),
		"video": templates.LongDesc(i18n.T(`
			Send an video file
			
			send an video file .mp4, 3gp supported by whatsapp, you need to provide the path to the file
			from there it will read and uploaded it to the whatsapp infra, after that it will 
			send to the user
		`)),
	}
	// ExampleDesc defines a map for all the available commands
	// it shows the current usage of each command all the available flags
	ExampleDesc = map[string]string{
		"text": templates.Examples(i18n.T(`
			# send a text message
			wabactl send text --to 9190000000 --message <string>
		`)),
		"image": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send image --to 9190000000 --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send image --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
		"document": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send document --to 9190000000 --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send document --to <wa_id with country code> --caption[optional] --url <file url>
		`)),
		"audio": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send audio --to 9190000000 --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send audio --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
		"video": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send video --to 9190000000 --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send video --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
	}
)
