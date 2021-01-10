package main

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
)

func init()  {
	// initialize config and other necessary packages
	_, err := config.New()
	if err != nil {
		log.Error(err)
	}
}

func main() {
	log.Info("started package successfully")
}
