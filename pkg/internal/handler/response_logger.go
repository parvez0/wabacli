package handler

import (
	"encoding/json"
	"fmt"
)

var logger = fmt.Println

func JsonResponse(msg interface{})  {
	var res map[string]interface{}
	if js, ok := msg.(string); ok {
		_ = json.Unmarshal([]byte(js), &res)
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			logger(js)
		}
		logger(string(b))
	} else {
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			logger(msg)
		}
		logger(string(b))
	}
}
