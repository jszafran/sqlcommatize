package main

import (
	"flag"
	sqlcommatize "github.com/jszafran/go-sql-commas"
	"golang.design/x/clipboard"
)

func main() {
	var forStrings bool
	var leadingCommas bool
	flag.BoolVar(&forStrings, "strings", false, "Wraps rows with single quotes (for SQL strings).")
	flag.BoolVar(&leadingCommas, "leading_commas", false, "Use leading commas for separating rows "+
		"(trailing commas used by default).")
	flag.Parse()

	cs := sqlcommatize.Trailing
	if leadingCommas {
		cs = sqlcommatize.Leading
	}

	rt := sqlcommatize.Number
	if forStrings {
		rt = sqlcommatize.String
	}

	inp := clipboard.Read(clipboard.FmtText)
	res := sqlcommatize.Commatize(inp, rt, cs)
	clipboard.Write(clipboard.FmtText, res)
}
