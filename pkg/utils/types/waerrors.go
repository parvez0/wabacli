package types

import "fmt"

type Meta struct {
	Version   string `json:"version"`
	APIStatus string `json:"api_status"`
}

type WhatsappError struct {
	Meta Meta `json:"meta"`
	Errors []struct {
		Code    int    `json:"code"`
		Title   string `json:"title"`
		Details string `json:"details"`
	} `json:"errors"`
}

func (w *WhatsappError) Error() string {
	return fmt.Sprintf("%+v", w.Errors)
}