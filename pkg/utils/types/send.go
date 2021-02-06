package types

type MessageType string

// All the available message type supported by Whatsapp
const (
	MessageTypeText MessageType = "text"
	MessageTypeImage MessageType = "image"
	MessageTypeAudio MessageType = "audio"
	MessageTypeVideo MessageType = "video"
	MessageTypeDocument MessageType = "document"
	MessageTypeTemplateText MessageType = "hsm"
	MessageTypeTemplateMedia MessageType = "template"
	MessageTypeContact MessageType = "contacts"
	MessageTypeLocation MessageType = "location"
	MessageTypeSticker MessageType = "sticker"
)

// MessageTypesMapping provides a mapping to message types
// which will be use to validate the user specified operation
var (
	MessageTypesMapping = map[MessageType]bool{
		MessageTypeText: true,
		MessageTypeImage: true,
		MessageTypeAudio: true,
		MessageTypeVideo: true,
		MessageTypeDocument: true,
		MessageTypeTemplateText: true,
		MessageTypeTemplateMedia: true,
		MessageTypeContact: true,
		MessageTypeLocation: true,
		MessageTypeSticker: true,
	}
)

type WAMessage struct {
	To            string `json:"to"`
	Type          string `json:"type"`
	RecipientType string `json:"recipient_type"`
	PreviewURL    bool   `json:"preview_url"`
	Text          struct {
		Body string `json:"body"`
	} `json:"text"`
	Image MediaMessage `json:"image"`
	Audio MediaMessage `json:"audio"`
	Video MediaMessage `json:"video"`

}

type Provider struct {
	Name string `json:"name"`
}

type MediaMessage struct {
	Provider Provider `json:"provider"`
	Caption string `json:"caption"`
	ID      string `json:"id"`
	Link    string `json:"link"`
}
