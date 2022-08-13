package main

import (
	sqlcomm "github.com/jszafran/go-sql-commas"
	"golang.design/x/clipboard"
)

func main() {
	clpText := sqlcomm.ReadClipboard()
	spl := sqlcomm.SplitContent(clpText)
	res := sqlcomm.WrapCommas(spl)
	clipboard.Write(sqlcomm.ClipboardFormat, res)
}
