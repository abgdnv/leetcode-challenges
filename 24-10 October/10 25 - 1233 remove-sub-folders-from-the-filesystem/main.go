package main

import (
	"regexp"
	"slices"
)

func removeSubfolders(folder []string) []string {
	slices.Sort(folder)
	result := make([]string, 1, len(folder))
	result[0] = folder[0]

	var currentFolder string
	for i := 1; i < len(folder); i++ {
		currentFolder = result[len(result)-1]
		if m, err := regexp.Match("^"+currentFolder+"/", []byte(folder[i])); err == nil && m {
			continue
		}
		result = append(result, folder[i])
	}
	return result
}
