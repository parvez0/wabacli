package handler

import (
	"encoding/json"
	"fmt"
)

var logger = fmt.Println

func JsonResponse(js string)  {
	var res map[string]interface{}
	_ = json.Unmarshal([]byte(js), &res)
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		logger(js)
	}
	logger(string(b))
}
