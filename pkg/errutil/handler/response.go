package handler

import "fmt"

var logger = fmt.Println

func JsonResponse(js string) {
	logger(js)
}
