package templates

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"os"
	"reflect"
)

var (
	ConfigClusterHeaders = []string{"Number", "Name", "Insecure", "Server"}
)

type TableWriter struct {
	Data interface{}
	TW *tablewriter.Table
}

// NewTableWriter initializes a new table writer object,
// all the settings has been predefined it will print a table
// without borders and left align content
func NewTableWriter(data interface{}) *TableWriter {
	
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	
	return &TableWriter{
		Data: data,
		TW: table,
	}
}

func (tw *TableWriter)ProcessData() {
	switch tw.Data.(type) {
	case []config.Cluster:
		tw.TW.SetHeader(ConfigClusterHeaders)
		clusters := tw.Data.([]config.Cluster)
		var data [][]string
		for _, v := range clusters {
			if v.Context == config.DefaultCurrentContext {
				continue
			}
			val := reflect.ValueOf(v)
			var row []string
			for _, h := range ConfigClusterHeaders {
				row = append(row, fmt.Sprintf("%v", val.FieldByName(h)))
			}
			data = append(data, row)
		}
		tw.TW.AppendBulk(data)
	case []string:
		clusters := tw.Data.([]string)
		tw.TW.Append(clusters)
	default:
		log.Panic("failed to process data for table writer")
	}
}

func (tw *TableWriter)Render()  {
	tw.TW.Render()
}