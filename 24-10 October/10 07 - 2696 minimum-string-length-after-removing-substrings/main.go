package main

import "strings"

func minLength(s string) int {
	l1 := len(s)
	l2 := 0
	for l1 != l2 {
		l1 = len(s)
		s = strings.Join(strings.Split(s, "AB"), "")
		s = strings.Join(strings.Split(s, "CD"), "")
		l2 = len(s)
	}
	return l1
}
