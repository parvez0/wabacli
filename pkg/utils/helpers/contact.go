package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/internal/request"
	"github.com/parvez0/wabacli/pkg/utils/types"
	"net/http"
)

// VerifyContact is a helper function to call the verify http request to verify the end user number
// if the user is not registered with whatsapp it will result in an error and exit with status code 1
func VerifyContact(clus *config.Cluster, to int, blocking bool, forceCheck bool) (*types.ContactResponse, error) {
	url := clus.Server + "/v1/contacts"
	reqBody := map[string]interface{}{
		"contacts": []string{
			fmt.Sprintf("+%d", to),
		},
		"force_check": forceCheck,
	}
	if blocking {
		reqBody["blocking"] = "wait"
	}
	client := request.NewClient(clus)
	opts := request.Options{
		Url:     url,
		Method:  http.MethodPost,
		Body:    reqBody,
	}
	req, err := client.NewRequest(opts)
	if err != nil {
		return  nil, err
	}
	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	byts, err := res.GetBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body - %s", err.Error())
	}
	if res.GetStatusCode() == 200 {
		log.Debug("request got successful with status code - ", res.GetStatusCode())
		var resp types.ContactResponse
		json.Unmarshal(byts, &resp)
		return &resp, nil
	}
	return nil, fmt.Errorf("failed to verify contact - statusCode: %d - Error: %v", res.GetStatusCode(), string(byts))
}