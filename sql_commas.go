package go_sql_commas

import (
	"fmt"
	xclp "golang.design/x/clipboard"
	"strings"
)

type CommaStyle string

const (
	Leading  CommaStyle = "leading"
	Trailing CommaStyle = "trailing"
)

type RowType string

const (
	String RowType = "string"
	Number RowType = "number"
)

type dataStore interface {
	Read() []byte
	Write(b []byte)
}

type Commatizer struct {
	store dataStore
}

func (c *Commatizer) Transform(rt RowType, cs CommaStyle) {
	rows := readRows(c.store.Read())
	if rt == String {
		rows = addSingleQuotes(rows)
	}
	withCommas := addCommas(rows, cs)
	c.store.Write(withCommas)
}

// clipboardStore is a data store based on OS clip-board.
type clipboardStore struct{}

// Read reads data from OS clip-board
func (cs *clipboardStore) Read() []byte {
	return xclp.Read(xclp.FmtText)
}

// Write writes b to OS clip-board
func (cs *clipboardStore) Write(b []byte) {
	xclp.Write(xclp.FmtText, b)
}

// fakeStore is an util store used for testing - stores data within struct field.
type fakeStore struct {
	data []byte
}

// Read reads data from fake store
func (fs *fakeStore) Read() []byte {
	return fs.data
}

// Write writes b to fake store
func (fs *fakeStore) Write(b []byte) {
	fs.data = b
}

// addCommas adds trailing/leading commas to given rows,
// so they can be later used within SQL IN clause.
// Returns result as slice of bytes.
func addCommas(rows []string, commaStyle CommaStyle) []byte {
	var r []byte

	switch commaStyle {
	case Trailing:
		r = []byte(strings.Join(rows, fmt.Sprintf(",%s", LineBreak)))
	case Leading:
		res := make([]string, len(rows))
		for i, row := range rows {
			if i == 0 {
				res[i] = fmt.Sprintf("%s", row)
				continue
			}
			res[i] = fmt.Sprintf(",%s", row)
		}
		r = []byte(strings.Join(res, LineBreak))
	}
	return r
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

func NewCommatizer() Commatizer {
	return Commatizer{store: &clipboardStore{}}
}

func testCommatizerWithData(b []byte) Commatizer {
	return Commatizer{store: &fakeStore{b}}
}
