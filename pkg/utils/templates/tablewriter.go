package templates

import (
	"github.com/olekukonko/tablewriter"
	"github.com/parvez0/wabacli/log"
	"os"
	"reflect"
)

const  (
	IgnoreFieldAuth = "Auth"
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

func (tw *TableWriter)WriteHeaders() {
	if reflect.ValueOf(tw.Data).Kind() != reflect.Struct {
		log.Panic("expected struct found ", reflect.TypeOf(tw.Data))
	}
	val := reflect.TypeOf(tw.Data)
	var keys []string
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Name == IgnoreFieldAuth {
			continue
		}
		keys = append(keys, val.Field(i).Name)
	}
	log.Debug("got headers, for table writer -", keys)
	tw.TW.SetHeader(keys)
}

func (tw *TableWriter)WriteData()  {
	tw.TW.Append([]string{"oiasjdfojsadfasdfsadfasdfasdfasdfasdfasdfasf", "localhost", "nemsa", "Insecure", "this"})
}

func (tw *TableWriter)Render()  {
	tw.TW.Render()
}