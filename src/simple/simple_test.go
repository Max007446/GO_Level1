package simpletest

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name  string
	input int
	want  []int
}{
	{name: "10", input: 10, want: []int{1, 2, 3, 5, 7}},
	{name: "20", input: 20, want: []int{1, 2, 3, 5, 7, 11, 13, 17, 19}},
	{name: "30", input: 30, want: []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	{name: "0", input: 0, want: []int{1}},
}

func TestSimple(t *testing.T) {
	for _, tc := range tests {
		got := Simple(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
