package handler

import (
	"encoding/json"
	"github.com/parvez0/wabacli/log"
)

var fatalError = log.Error

// FatalLog will use log.Error() to exit after
// logging the error
func FatalError(err error)  {
	if err != nil {
		fatalError("error: ", err)
	}
}

func FatalJsonError(err error)  {
	if err != nil {
		js, err := json.MarshalIndent(err, "", "  ")
		if err != nil {
			fatalError("error: ", err)
		}
		fatalError(string(js))
	}
}