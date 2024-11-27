package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	testCases := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"My name is Haley", "My Haley", true},
		{"of", "A lot of words", false},
		{"Eating right now", "Eating", true},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintln(tC.s1, tC.s2, tC.want), func(t *testing.T) {
			value := areSentencesSimilar(tC.s1, tC.s2)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
