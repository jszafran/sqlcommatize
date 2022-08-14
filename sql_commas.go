package go_sql_commas

import (
	"fmt"
	"golang.design/x/clipboard"
	"strings"
)

const ClipboardFormat = clipboard.FmtText

type Clipboard interface {
	ReadFrom() string
	WriteTo(b []byte)
}

type FakeClipboard struct {
	Data string
}

func (fc *FakeClipboard) ReadFrom() string {
	return fc.Data
}

func (fc *FakeClipboard) WriteTo(b []byte) {
	fc.Data = string(b)
}

type SystemClipboard struct{}

func (sc *SystemClipboard) ReadFrom() string {
	return string(clipboard.Read(ClipboardFormat))
}

func (sc *SystemClipboard) WriteTo(b []byte) {
	clipboard.Write(ClipboardFormat, b)
}

// addTrailingComma adds trailing comma to records (except for the last one)
// so they can be later used within SQL IN clause. Returns slice of bytes.
func addTrailingComma(rows []string) []byte {
	return []byte(strings.Join(rows, ",\n"))
}

// addSingleQuotes wraps all records with single quotes. If record contains a single quote,
// it gets replaced with double single quotes
func addSingleQuotes(rows []string) []string {
	r := make([]string, len(rows))
	for i, rec := range rows {
		r[i] = fmt.Sprintf("'%s'", strings.Replace(rec, "'", "''", -1))
	}
	return r
}

// readRows reads content from a clipboard (treating it as a text)
// and splits them into records
func readRows(clp Clipboard) []string {
	txt := clp.ReadFrom()
	return strings.Split(txt, "\n")
}

// HandleNumbers processes and treats clipboard content as if it's SQL numeric values.
// Comma is appended to every row (except for the last one), so you can use it later
// in a SQL IN clause like that:
//		SELECT * FROM FOO
// 		WHERE 1=1
//		AND ID IN (
//			1,
//			2,
//			3
//		)
func HandleNumbers(clp Clipboard) {
	rows := readRows(clp)
	processedRows := addTrailingComma(rows)
	clp.WriteTo(processedRows)
}

// HandleStrings processes and treats clipboard content as if it's SQL string value.
// Each row is wrapped with single quotes and & appended with comma (comma skipped only for the last one),
// so it can be later used in a SQL IN clause like that:
//		SELECT * FROM FOO
// 		WHERE 1=1
//		AND NAME IN (
//			'Foo',
//			'Bar',
//			'Baz'
//		)
func HandleStrings(clp Clipboard) {
	rows := readRows(clp)
	quotedRows := addSingleQuotes(rows)
	processedRows := addTrailingComma(quotedRows)
	clp.WriteTo(processedRows)
}
