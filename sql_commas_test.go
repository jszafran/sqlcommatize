package go_sql_commas

import (
	"fmt"
	"testing"
)

type testCase struct {
	input         string
	want          string
	leadingCommas bool
}

// FakeClipboard is a helper struct for testing - implements Clipboard interface.
// Allows to imitate the behaviour of the clipboard and check its value within tests.
type FakeClipboard struct {
	Data string
}

func (fc *FakeClipboard) ReadFrom() string {
	return fc.Data
}

func (fc *FakeClipboard) WriteTo(b []byte) {
	fc.Data = string(b)
}
func TestHandleNumbers(t *testing.T) {
	t.Parallel()
	cases := []testCase{
		{"1", "1", false},
		{"1", "1", true},
		{fmt.Sprintf("1%s2", LineBreak), fmt.Sprintf("1,%s2", LineBreak), false},
		{fmt.Sprintf("1%s2", LineBreak), fmt.Sprintf("1%s,2", LineBreak), true},
		{fmt.Sprintf("1%s2%s3", LineBreak, LineBreak), fmt.Sprintf("1,%s2,%s3", LineBreak, LineBreak), false},
		{fmt.Sprintf("1%s2%s3", LineBreak, LineBreak), fmt.Sprintf("1%s,2%s,3", LineBreak, LineBreak), true},
		//Excel adding extra line break at the end edge case
		{fmt.Sprintf("1%s2%s", LineBreak, LineBreak), fmt.Sprintf("1,%s2", LineBreak), false},
		{fmt.Sprintf("1%s2%s", LineBreak, LineBreak), fmt.Sprintf("1%s,2", LineBreak), true},
	}
	for _, c := range cases {
		clp := FakeClipboard{c.input}
		HandleNumbers(&clp, c.leadingCommas)
		got := clp.Data
		if got != c.want {
			t.Fatalf("Want %s, got %s", c.want, got)
		}
	}
}

func TestHandleStrings(t *testing.T) {
	t.Parallel()
	cases := []testCase{
		{"a", "'a'", false},
		{"a", "'a'", false},
		{fmt.Sprintf("a%sb", LineBreak), fmt.Sprintf("'a',%s'b'", LineBreak), false},
		{fmt.Sprintf("a%sb", LineBreak), fmt.Sprintf("'a'%s,'b'", LineBreak), true},
		{fmt.Sprintf("a%sb%sc", LineBreak, LineBreak), fmt.Sprintf("'a',%s'b',%s'c'", LineBreak, LineBreak), false},
		{fmt.Sprintf("a%sb%sc", LineBreak, LineBreak), fmt.Sprintf("'a'%s,'b'%s,'c'", LineBreak, LineBreak), true},
		{fmt.Sprintf("'a'%sb%s'c''", LineBreak, LineBreak), fmt.Sprintf("'''a''',%s'b',%s'''c'''''", LineBreak, LineBreak), false},
		{fmt.Sprintf("'a'%sb%s'c''", LineBreak, LineBreak), fmt.Sprintf("'''a'''%s,'b'%s,'''c'''''", LineBreak, LineBreak), true},
		// Excel adding extra line break at the end edge case
		{fmt.Sprintf("1'%s'2%s", LineBreak, LineBreak), fmt.Sprintf("'1''',%s'''2'", LineBreak), false},
		{fmt.Sprintf("1'%s'2%s", LineBreak, LineBreak), fmt.Sprintf("'1'''%s,'''2'", LineBreak), true},
	}
	for _, c := range cases {
		clp := FakeClipboard{c.input}
		HandleStrings(&clp, c.leadingCommas)
		got := clp.Data
		if got != c.want {
			t.Fatalf("Want %s, got %s", c.want, got)
		}
	}
}
