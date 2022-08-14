package go_sql_commas

import "testing"

type testCase struct {
	input string
	want  string
}

// FakeClipboard is a helper struct for testing.
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
		{"1", "1"},
		{"1\n2", "1,\n2"},
		{"1\n2\n3\n4", "1,\n2,\n3,\n4"},
	}
	for _, c := range cases {
		clp := FakeClipboard{c.input}
		HandleNumbers(&clp)
		got := clp.Data
		if got != c.want {
			t.Fatalf("Want %s, got %s", c.want, got)
		}
	}
}

func TestHandleStrings(t *testing.T) {
	t.Parallel()
	cases := []testCase{
		{"a", "'a'"},
		{"a\nb", "'a',\n'b'"},
		{"a\nb \nc", "'a',\n'b ',\n'c'"},
		{"'a'\nb\n'c''", "'''a''',\n'b',\n'''c'''''"},
	}
	for _, c := range cases {
		clp := FakeClipboard{c.input}
		HandleStrings(&clp)
		got := clp.Data
		if got != c.want {
			t.Fatalf("Want %s, got %s", c.want, got)
		}
	}
}
