package request

import (
	"github.com/parvez0/wabacli/config"
	"net/http"
	"testing"
)

var cluster = config.Cluster{
	Auth:     "",
	Server:   "https://ipinfo.io",
	Name:     "",
	Number:   0,
	Insecure: true,
	Context:  "",
}

var client = NewClient(&cluster)

func TestRequestWrapper(t *testing.T)  {
	options := Options{
		Url:     "/",
		Method:  http.MethodGet,
	}
	req, err := client.NewRequest(options)
	if err != nil {
		t.Fatal("GetRequestFailed: " + err.Error())
	}
	res, err := req.Send()
	if err != nil {
		t.Fatal("FailedApiCall: " + err.Error())
	}
	buf, err := res.GetBody()
	if err != nil {
		t.Fatal("FailedToReadResponseBody: " + err.Error())
	}
	if res.GetStatusCode() != http.StatusOK {
		t.Fatalf("FailedWithStatusCode: status_code: %d, body: %s", res.GetStatusCode(), string(buf))
	}
	t.Logf("Api call success, response_code: %d, body: %s", res.GetStatusCode(), string(buf))
}
