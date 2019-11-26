package yogatable

import (
	"github.com/olekukonko/tablewriter"
	"io"
)

// PCSTable 封装 tablewriter.Table
type YogaTable struct {
	*tablewriter.Table
}


// NewTable 预设了一些配置
func NewTable(wt io.Writer) YogaTable {
	tb := tablewriter.NewWriter(wt)
	tb.SetAutoWrapText(false)
	tb.SetBorder(false)
	tb.SetHeaderLine(false)
	tb.SetColumnSeparator("")
	return YogaTable{tb}
}

