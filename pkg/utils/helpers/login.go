package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/badrequest"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	handler2 "github.com/parvez0/wabacli/pkg/internal/handler"
	"github.com/parvez0/wabacli/pkg/internal/request"
	"net/http"
)

// Login performs login into the whatsapp account and generates
// a long term auth token, if new password is provided the
// existing will be replaced by the new password
func Login(c *config.Cluster, pwd string, np string) string {
	//Creating a new request client with cluster information
	client := request.NewClient(c)
	reqBody := make(map[string]string)
	log.Debug("initiating the login procedure")
	if np != "" {
		log.Debug("resetting the password")
		reqBody["new_password"] = np
	}
	// Headers for specifying the content type of the body
	// and also for giving the base 64 encoded credentials
	// in username:password format
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + pwd)))
	//Creating object the options for the request call by
	// default url will be switch to https://localhost if the
	// server is not running on
	opts := request.Options{
		Url:     "/v1/users/login",
		Method:  http.MethodPost,
		Body:    reqBody,
		Headers: headers,
	}
	log.Debug(fmt.Sprintf("login request options: %+v", opts))
	req, err := client.NewRequest(opts)
	handler.FatalError(err)
	// Making the api call
	res, err := req.Send()
	handler.FatalError(err)
	log.Debug("login request successful with status_code:", res.GetStatusCode())
	// After getting the response, if the status is 200 then user
	// is able to login and the auth is generated
	if res.GetStatusCode() != http.StatusOK {
		buf, _ := res.GetBody()
		handler.FatalJsonError(&badrequest.BadRequest{
			Code:        res.GetStatusCode(),
			Title:       "Login request failed",
			Description: string(buf),
		})
	}
	buf, err := res.GetBody()
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to process the login response: %v", err))
	}
	var loginResp handler2.LoginResponse
	err = json.Unmarshal(buf, &loginResp)
	handler.FatalError(err)
	js, err := json.MarshalIndent(loginResp, "", "  ")
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to parse response: %s", string(buf)))
	}
	// logging the json response to the screen if the
	// auth token generated successfully
	handler2.JsonResponse(string(js))
	auth := loginResp.Users[0].Token
	return auth
}
