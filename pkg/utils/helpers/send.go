package helpers

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/internal/request"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"net/http"
)

// SendMessage is a helper function to make a http call the whatsapp infra
func SendMessage(clus *config.Cluster, msg *map[string]interface{}) ([]byte, error) {
	url := clus.Server + "/v1/messages"
	client := request.NewClient(clus)
	opts := request.Options{
		Url:     url,
		Method:  http.MethodPost,
		Body:    msg,
	}
	req, err := client.NewRequest(opts)
	if err != nil {
		return nil, err
	}
	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if res.GetStatusCode() == 201 || res.GetStatusCode() == 200 {
		log.Debug("request got successful with status code - ", res.GetStatusCode())
	}
	return res.GetBody()
}

// UploadMedia provides a helper function to upload the file to whatsapp
func UploadMedia(clus *config.Cluster, file *types.Media) ([]byte, error) {
	url := clus.Server + "/v1/media"
	client := request.NewClient(clus)
	opts := request.Options{
		Url:     url,
		Method:  http.MethodPost,
		Headers: client.AddHeader("Content-Type", file.GetMimeType()),
		Body:    file.Data,
	}
	req, err := client.NewRequest(opts)
	if err != nil {
		return nil, err
	}
	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if res.GetStatusCode() == 201 || res.GetStatusCode() == 200 {
		log.Debug("request got successful with status code - ", res.GetStatusCode())
	}
	return res.GetBody()
}