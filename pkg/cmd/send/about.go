package send

import (
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	ShortDesc = map[string]string {
		"text": i18n.T("Send a text message"),
		"image": i18n.T("Send an image file"),
		"doc": i18n.T("Send a document"),
		"audio": i18n.T("Send an audio file"),
		"video": i18n.T("Send a video file"),
	}
	LongDesc = map[string]string{
		"text": templates.LongDesc(i18n.T(`
			Send a text message
			
			send allows you to send a text message to a user,  it sends the message
			to the whatsapp infrastructure from there a job will be created which
			will take care of sending the message to the user
		`)),
		"image": templates.LongDesc(i18n.T(`
			Send an image file
			
			send a jpeg or a png file to the user, you need provide the path to the file
			or you can directly give the url, from there it will read and uploaded to the
			whatsapp infra, after that it will send to the user
		`)),
		"doc": templates.LongDesc(i18n.T(`
			Send an document
			
			send a document supported by whatsapp, you need to provide the path to the file 
			or you can directly give the url, from there it will read and uploaded to the
			whatsapp infra, after that it will send to the user
		`)),
		"audio": templates.LongDesc(i18n.T(`
			Send an audio file
			
			send an audio file (.mp3 or .mp4) supported by whatsapp, you need provide the path to the file
			from there it will read and uploaded to the whatsapp infra, after that it will 
			send to the user
		`)),
		"video": templates.LongDesc(i18n.T(`
			Send an video file
			
			send an video file .mp4, 3gp supported by whatsapp, you need provide the path to the file
			from there it will read and uploaded to the whatsapp infra, after that it will 
			send to the user
		`)),
	}
	ExampleDesc = map[string]string{
		"text": templates.Examples(i18n.T(`
			# send a text message
			wabactl send text --to <wa_id with country code> --message <string>
		`)),
		"image": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send image --to <wa_id with country code> --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send image --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
		"doc": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send document --to <wa_id with country code> --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send document --to <wa_id with country code> --caption[optional] --url <file url>
		`)),
		"audio": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send audio --to <wa_id with country code> --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send audio --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
		"video": templates.Examples(i18n.T(`
			# send an video with a caption
			wabactl send video --to <wa_id with country code> --caption[optional] <string> --path <path/to/file>
			# send an video using url
			wabactl send video --to <wa_id with country code> --caption[optional] <string> --url <image url>
		`)),
	}
)
