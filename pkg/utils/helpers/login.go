package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/cmd/context"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/internal/request"
	"net/http"
)

func Login(ap *context.AddOptions, ) (string, error) {
	client := request.NewClient(ap.Cluster)
	reqBody := make(map[string]string)
	log.Debug("initiating the login procedure")
	if ap.Reset {
		log.Debug("resetting the password")
		reqBody["new_password"] = ap.NewPassword
	}
	opts := request.Options{
		Url:     "/v1/users/login",
		Method:  http.MethodPost,
		Body:    reqBody,
	}
	log.Debug(fmt.Sprintf("login request options: %+v", opts))
	req, err := client.NewRequest(opts)
	handler.FatalError(err)
	res, err := req.Send()
	handler.FatalError(err)
	log.Debug("login request successful with status_code:", res.GetStatusCode())
	if res.GetStatusCode() != http.StatusOK {
		handler.FatalJsonError(&badrequest.BadRequest{
			Code:        res.GetStatusCode(),
			Title:       "Login request failed",
			Description: err.Error(),
		})
	}
	buf, err := res.GetBody()
	handler.FatalError(fmt.Errorf("failed to process the login response: %s", err.Error()))
	js, err := json.MarshalIndent(buf, "", "  ")
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to parse response: %s", string(buf)))
	}
	handler.JsonResponse(string(js))
	return "", nil
}
