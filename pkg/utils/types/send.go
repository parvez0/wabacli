package types

import "fmt"

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
	To            string `json:"to" validate:"required"`
	Type          string `json:"type" validate:"required"`
	RecipientType string `json:"recipient_type"`
	PreviewURL    bool   `json:"preview_url"`
	Caption string `json:"caption"`
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

type MediaResponse struct {
	Media []struct {
		ID string `json:"id"`
	} `json:"media"`
	Meta Meta `json:"meta"`
}

func (m *MediaResponse) GetId() (string, error) {
	if len(m.Media) > 0 {
		return m.Media[0].ID, nil
	}
	return "", fmt.Errorf("media id is not generated - %+v", m)
}
