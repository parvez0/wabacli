package main

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/cmd"
)

func init()  {
	// initialize config and other necessary packages
	_, err := config.GetConfig()
	if err != nil {
		log.Error(err)
	}
}

func main() {
	cmd.Execute()
}
