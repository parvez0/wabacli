package context

import "github.com/parvez0/wabacli/config"

type GetOptions struct {
	Config *config.Configuration
	Selector string
}

func (o *GetOptions)Validate() {
	if o.Selector == ""{
		o.Selector
	}
}
