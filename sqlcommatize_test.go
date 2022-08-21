package sqlcommatize

import (
	"fmt"
	"reflect"
	"testing"
)

const LB = LineBreak

func TestCommatizer_Transform(t *testing.T) {
	type testCase struct {
		input string
		want  string
	}

	runCases := func(tc []testCase, rt RowType, cs CommaStyle) {
		for _, c := range tc {
			got := Commatize([]byte(c.input), rt, cs)
			want := []byte(c.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("Got %v, want %v", got, want)
			}
		}
	}

	// numbers
	cases := []testCase{
		{"1", "1"},
		{fmt.Sprintf("1%s2", LB), fmt.Sprintf("1,%s2", LB)},
		{fmt.Sprintf("1%s2%s3", LB, LB), fmt.Sprintf("1,%s2,%s3", LB, LB)},
		// Excel adding extra line break at the end edge case
		{fmt.Sprintf("1%s2%s", LB, LB), fmt.Sprintf("1,%s2", LB)},
	}
	runCases(cases, Number, Trailing)

	// strings
	cases = []testCase{
		{"a", "'a'"},
		{fmt.Sprintf("a%sb", LB), fmt.Sprintf("'a',%s'b'", LB)},
		{fmt.Sprintf("a%sb%sc", LB, LB), fmt.Sprintf("'a',%s'b',%s'c'", LB, LB)},
		{fmt.Sprintf("'a'%sb%s'c''", LB, LB), fmt.Sprintf("'''a''',%s'b',%s'''c'''''", LB, LB)},
		// Excel adding extra line break at the end edge case
		{fmt.Sprintf("1'%s'2%s", LB, LB), fmt.Sprintf("'1''',%s'''2'", LB)},
	}
	runCases(cases, String, Trailing)

	// leading commas
	cases = []testCase{
		{"a", "'a'"},
		{fmt.Sprintf("a%sb", LB), fmt.Sprintf("'a'%s,'b'", LB)},
		{fmt.Sprintf("a%sb%sc", LB, LB), fmt.Sprintf("'a'%s,'b'%s,'c'", LB, LB)},
	}
	runCases(cases, String, Leading)

}
