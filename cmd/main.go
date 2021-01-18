package main

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/cmd"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
)

func init()  {
	// initialize config and other necessary packages
	_, err := config.GetConfig()
	if err != nil {
		handler.FatalError(err)
	}
}

func main() {
	cmd.Execute()
}
