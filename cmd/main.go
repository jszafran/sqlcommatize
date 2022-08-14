package main

import (
	"flag"
	go_sql_commas "github.com/jszafran/go-sql-commas"
)

func main() {
	var forStrings bool
	flag.BoolVar(&forStrings, "strings", false, "Wraps output with single quotes (SQL strings).")
	flag.Parse()

	clp := go_sql_commas.SystemClipboard{}

	if forStrings {
		go_sql_commas.HandleStrings(&clp)
	} else {
		go_sql_commas.HandleNumbers(&clp)
	}
}
