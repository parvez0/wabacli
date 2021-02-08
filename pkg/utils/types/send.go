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

// WAMessage defines all the message formats and it's
// values which will further be process by intermediate
// functions to send the request to whatsapp
type WAMessage struct {
	To            string `json:"to" validate:"required"`
	Type          string `json:"type" validate:"required"`
	RecipientType string `json:"recipient_type"`
	PreviewURL    bool   `json:"preview_url,omitempty"`
	Caption string `json:"caption,omitempty"`
	Text Text `json:"text,omitempty"`
	Image MediaMessage `json:"image,omitempty"`
	Audio MediaMessage `json:"audio,omitempty"`
	Video MediaMessage `json:"video,omitempty"`
	Document MediaMessage `json:"document,omitempty"`
}

type Text struct {
	Body string `json:"body"`
}

type Provider struct {
	Name string `json:"name"`
}

type MediaMessage struct {
	Provider Provider `json:"provider"`
	Caption string `json:"caption"`
	ID      string `json:"id,omitempty"`
	Link    string `json:"link,omitempty"`
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

type MessageResponse struct {
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
	Meta Meta
}

func (m *MessageResponse) GetId() (string, error) {
	if len(m.Messages) > 0 {
		return m.Messages[0].ID, nil
	}
	return "", fmt.Errorf("message id is not generated - %+v", m)
}
