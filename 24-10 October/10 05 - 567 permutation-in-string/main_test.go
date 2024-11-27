package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		s1, s2 string
		want   bool
	}{
		{"abc", "bbbca", true},
		{"ab", "eidbaooo", true},
		{"adc", "dcda", true},
		{"ab", "eidboaoo", false},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintln(tC.s1, tC.s2, "want", tC.want), func(t *testing.T) {
			value := checkInclusion(tC.s1, tC.s2)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
