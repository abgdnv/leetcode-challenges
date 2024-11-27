package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{"ABFCACDB", 2},
		{"ACBBD", 5},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintln(tC), func(t *testing.T) {
			value := minLength(tC.s)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
