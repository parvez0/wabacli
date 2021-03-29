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
	CurrentContext string
	TW *tablewriter.Table
}

// NewTableWriter initializes a new table writer object,
// all the settings has been predefined it will print a table
// without borders and left align content
func NewTableWriter(data interface{}, context string) *TableWriter {
	
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
		CurrentContext: context,
	}
}

func (tw *TableWriter) ProcessData() {
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
			var colRow []tablewriter.Colors
			for _, h := range ConfigClusterHeaders {
				fVal := fmt.Sprintf("%v", val.FieldByName(h))
				if h == "Number" {
					fVal = fmt.Sprintf("%s******%s", fVal[0:2], fVal[len(fVal)-3:])
				}
				row = append(row, fVal)
				if v.Context == tw.CurrentContext {
					colRow = append(colRow, tablewriter.Colors{tablewriter.Normal, tablewriter.BgWhiteColor, tablewriter.FgBlackColor})
				}
			}
			if v.Context == tw.CurrentContext {
				tw.TW.Rich(row, colRow)
				continue
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