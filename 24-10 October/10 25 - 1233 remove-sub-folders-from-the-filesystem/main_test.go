package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		folder []string
		want   []string
	}{
		{
			[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			[]string{"/a", "/c/d", "/c/f"},
		},
		{
			[]string{"/a", "/a/b/c", "/a/b/d"},
			[]string{"/a"},
		},
		{
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := removeSubfolders(tC.folder)
			if !isEqual(value, tC.want) {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}

func isEqual(s1, s2 []string) bool {
	slices.Sort(s1)
	slices.Sort(s2)
	return reflect.DeepEqual(s1, s2)
}
