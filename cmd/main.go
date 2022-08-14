package main

import (
	"flag"
	go_sql_commas "github.com/jszafran/go-sql-commas"
)

func main() {
	var wrapStrings bool
	flag.BoolVar(&wrapStrings, "strings", false, "Wraps output with single quotes (SQL strings).")
	flag.Parse()

	clp := go_sql_commas.SystemClipboard{}

	var wrapped []byte
	if wrapStrings {
		wrapped = go_sql_commas.HandleStrings(clp)
	} else {
		wrapped = go_sql_commas.HandleNumbers(clp)
	}

	go_sql_commas.PasteToClipboard(wrapped)
}
