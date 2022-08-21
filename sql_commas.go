package go_sql_commas

import (
	"fmt"
	"strings"
)

type CommaStyle string

// Trailing comma style:
//	1,
//	2,
//	3
// Leading comma style:
//	1
//	,2
//	,3
const (
	Leading  CommaStyle = "leading"
	Trailing CommaStyle = "trailing"
)

type RowType string

// String row type means rows will be used as SQL strings
// - they need to be wrapped with single quotes.
// Number row type means rows will be used as SQL numeric values
// - no single quotes required.
const (
	String RowType = "string"
	Number RowType = "number"
)

// Commatize splits b input by OS break line character (\r\n or \n),
// wraps them optionally with single quotes (for SQL strings)
// and delimits with commas (applying trailing/leading comma styling based on cs argument).
// Returns result as a slice of bytes.
func Commatize(b []byte, rt RowType, cs CommaStyle) []byte {
	rows := readRows(b)
	if rt == String {
		rows = addSingleQuotes(rows)
	}
	return addCommas(rows, cs)
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
