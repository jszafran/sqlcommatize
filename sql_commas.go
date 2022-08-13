package go_sql_commas

import (
	"golang.design/x/clipboard"
	"strings"
)

const ClipboardFormat = clipboard.FmtText

func ReadClipboard() string {
	clpb := clipboard.Read(ClipboardFormat)
	return string(clpb)
}

func SplitContent(s string) []string {
	// start with most naive implementation
	return strings.Split(s, "\n")
}

func WrapCommas(recs []string) []byte {
	return []byte(strings.Join(recs, ",\n"))
}
