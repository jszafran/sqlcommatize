package go_sql_commas

import (
	"fmt"
	"golang.design/x/clipboard"
	"strings"
)

const ClipboardFormat = clipboard.FmtText

type Clipboard interface {
	Read() string
}

type FakeClipboard struct {
	Data string
}

func (fc FakeClipboard) Read() string {
	return fc.Data
}

type SystemClipboard struct{}

func (sc SystemClipboard) Read() string {
	return string(clipboard.Read(ClipboardFormat))
}

// splitContent split clipboard string into records based on separator character
func splitContent(s string) []string {
	// start with most naive implementation - assume \n is the separator
	return strings.Split(s, "\n")
}

// addTrailingComma adds trailing comma to records (except for the last one)
// so they can be later used within SQL IN clause. Returns slice of bytes.
func addTrailingComma(recs []string) []byte {
	return []byte(strings.Join(recs, ",\n"))
}

// addSingleQuotes wraps all records with single quotes. If record contains a single quote,
// it gets replaced with double single quotes
func addSingleQuotes(recs []string) []string {
	r := make([]string, len(recs))
	for i, rec := range recs {
		r[i] = fmt.Sprintf("'%s'", strings.Replace(rec, "'", "''", -1))
	}
	return r
}

func HandleNumbers(clp Clipboard) []byte {
	d := clp.Read()
	recs := splitContent(d)
	return addTrailingComma(recs)
}

func HandleStrings(clp Clipboard) []byte {
	d := clp.Read()
	recs := splitContent(d)
	quoted := addSingleQuotes(recs)
	return addTrailingComma(quoted)
}

func PasteToClipboard(b []byte) {
	clipboard.Write(ClipboardFormat, b)
}
