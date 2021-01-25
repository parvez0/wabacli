package send

import (
	"github.com/parvez0/wabacli/pkg/utils/templates"
	"k8s.io/kubectl/pkg/util/i18n"
)

var (
	sendLong = templates.LongDesc(i18n.T(`
	Send message to user
	
	You can send text, media videos to the user. Just provide the user id and message to send
	`))
	sendExample = templates.Examples(i18n.T(`
	# send text message
	wabactl send text 		
	 `))
)
