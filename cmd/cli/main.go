package main

import (
	"flag"
	go_sql_commas "github.com/jszafran/go-sql-commas"
)

func main() {
	var forStrings bool
	var leadingCommas bool
	flag.BoolVar(&forStrings, "strings", false, "Wraps rows with single quotes (for SQL strings).")
	flag.BoolVar(&leadingCommas, "leading_commas", false, "Use leading commas for separating rows "+
		"(trailing commas used by default).")
	flag.Parse()

	cs := go_sql_commas.Trailing
	if leadingCommas {
		cs = go_sql_commas.Leading
	}

	rt := go_sql_commas.Number
	if forStrings {
		rt = go_sql_commas.String
	}

	cmtz := go_sql_commas.NewCommatizer()
	cmtz.Transform(rt, cs)
}
