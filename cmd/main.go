package main

import (
	"flag"
	go_sql_commas "github.com/jszafran/go-sql-commas"
)

func main() {
	var forStrings bool
	var leadingCommas bool
	flag.BoolVar(&forStrings, "strings", false, "Wraps rows with single quotes (SQL strings).")
	flag.BoolVar(&leadingCommas, "leading_commas", false, "Use leading commas for separating rows "+
		"(trailing commas are used by default.)")
	flag.Parse()

	clp := go_sql_commas.SystemClipboard{}

	if forStrings {
		go_sql_commas.HandleStrings(&clp, leadingCommas)
	} else {
		go_sql_commas.HandleNumbers(&clp, leadingCommas)
	}
}
