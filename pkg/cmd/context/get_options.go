package context

import (
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/pkg/utils/templates"
)

type GetOptions struct {
	Config *config.Configuration
	Selector string
	All bool
	Headers []string
}

func (o *GetOptions)Complete() {
	if o.Selector == ""{
		o.All = true
	}
}

func (o *GetOptions)Run() error {
	if o.All {
		tw := templates.NewTableWriter(o.Config.CurrentCluster)
		tw.WriteHeaders()
		tw.WriteData()
		tw.Render()
	}
	return nil
}
