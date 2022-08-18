package go_sql_commas

import (
	"fmt"
	xclp "golang.design/x/clipboard"
	"strings"
)

const ClipboardFormat = xclp.FmtText

// clipboardData defines a simple interface for interacting with clipboard
type clipboardData interface {
	Read() []byte
	Write(b []byte)
}

// systemClipboardData is a wrapper around actual operating system clip-board.
type systemClipboardData struct{}

// Read reads data from OS clip-board and returns it as slice of bytes.
func (scd *systemClipboardData) Read() []byte {
	return xclp.Read(ClipboardFormat)
}

// Write writes given data to OS clip-board.
func (scd *systemClipboardData) Write(b []byte) {
	xclp.Write(ClipboardFormat, b)
}

// fakeClipboardData is a helper struct for testing
// (to avoid reading from/writing to actual OS clip-board).
type fakeClipboardData struct {
	d []byte
}

// Read reads fake clip-board data.
func (fcd *fakeClipboardData) Read() []byte {
	return fcd.d
}

// Write writes given data to fake clip-board.
func (fcd *fakeClipboardData) Write(b []byte) {
	fcd.d = b
}

func SystemClipboard() *clipboard {
	return &clipboard{&systemClipboardData{}}
}

func fakeClipboard() *clipboard {
	return &clipboard{&fakeClipboardData{}}
}

type clipboard struct {
	data clipboardData
}

func (c *clipboard) Commatize(addQuotes bool, leadingCommas bool) {
	rows := readRows(c.data.Read())
	if addQuotes {
		rows = addSingleQuotes(rows)
	}
	c.data.Write(addCommas(rows, leadingCommas))
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
func readRows(b []byte) []string {
	s := string(b)
	// handle Excel edge case
	// it adds additional line break at the end of the copied content
	if len(s) >= len(LineBreak) && s[len(s)-len(LineBreak):] == LineBreak {
		s = s[:len(s)-len(LineBreak)]
	}
	return strings.Split(s, LineBreak)
}
