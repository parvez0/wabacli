package helpers

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/internal/request"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"net/http"
)

//func SendMessage(clus *config.Cluster, msg *types.WAMessage) ([]byte, error) {
//	url := clus.Server + "/v1/messages"
//	bearer := fmt.Sprintf("Bearer %s", clus.Auth)
//	byts, err := json.Marshal(msg)
//	if err != nil {
//		return nil, err
//	}
//	log.Debug("request json body generated - ", string(byts))
//	client := request.NewClient(clus)
//	buff := bytes.NewBuffer(byts)
//	req, err := http.NewRequest(http.MethodPost, url, buff)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set(HeaderAuthorization, bearer)
//	res, err := client.(req)
//	if err != nil {
//		return nil, err
//	}
//	byts, err = ioutil.ReadAll(res.Body)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read response: %s", err.Error())
//	}
//	logger.Debug("got response for message - from - ", api.Number, " - to - ", msg.To, " - res-body - ", string(byts))
//	if res.StatusCode != 200 || res.StatusCode != 201 {
//		var waErr types.WhatsappError
//		json.Unmarshal(byts, &waErr)
//		return nil, &waErr
//	}
//	return byts, err
//}


func SendMessage(clus *config.Cluster, msg *types.WAMessage) ([]byte, error) {
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