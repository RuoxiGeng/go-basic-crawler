package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	} {
		//Normal cases
		{"abcabcbb", 3},
		{"bbbbbbb", 1},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"abcdefj", 7},

		//Chinese
		{"网网王王王万", 2},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubstr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; " +
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B)  {
	s := "网网王王王万"
	for i := 0; i <13; i++ {
		s = s + s
	}
	ans := 3

	b.Logf("Len(s)=%d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubstr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d", actual, s, ans)
		}
	}
}