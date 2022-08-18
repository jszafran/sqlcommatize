package go_sql_commas

import (
	"fmt"
	"golang.design/x/clipboard"
	"strings"
)

type CommaStyle string

const ClipboardFormat = clipboard.FmtText

// Clipboard defines a simple interface for interacting with clipboard
type Clipboard interface {
	ReadFrom() []byte
	WriteTo(b []byte)
}

// SystemClipboard is a wrapper around golang.design/x/clipboard.
type SystemClipboard struct{}

func (sc *SystemClipboard) ReadFrom() []byte {
	return clipboard.Read(ClipboardFormat)
}

func (sc *SystemClipboard) WriteTo(b []byte) {
	clipboard.Write(ClipboardFormat, b)
}

// addCommas adds trailing/leading commas to given rows,
// so they can be later used within SQL IN clause.
// Returns result as slice of bytes.
func addCommas(rows []string, leadingCommas bool) []byte {
	// trailing commas style
	if !leadingCommas {
		return []byte(strings.Join(rows, fmt.Sprintf(",%s", LineBreak)))
	}

	// leading commas style
	res := make([]string, len(rows))
	for i, row := range rows {
		if i == 0 {
			res[i] = fmt.Sprintf("%s", row)
			continue
		}
		res[i] = fmt.Sprintf(",%s", row)
	}
	return []byte(strings.Join(res, LineBreak))
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
	txt := string(clp.ReadFrom())
	// handle Excel edge case
	// it adds additional line break at the end of the copied content
	if len(txt) >= len(LineBreak) && txt[len(txt)-len(LineBreak):] == LineBreak {
		txt = txt[:len(txt)-len(LineBreak)]
	}
	return strings.Split(txt, LineBreak)
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
func HandleNumbers(clp Clipboard, leadingCommas bool) {
	rows := readRows(clp)
	processedRows := addCommas(rows, leadingCommas)
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
func HandleStrings(clp Clipboard, leadingCommas bool) {
	rows := readRows(clp)
	quotedRows := addSingleQuotes(rows)
	processedRows := addCommas(quotedRows, leadingCommas)
	clp.WriteTo(processedRows)
}
