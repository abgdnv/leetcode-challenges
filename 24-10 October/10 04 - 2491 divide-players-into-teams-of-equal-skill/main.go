package main

import (
	"slices"
)

func dividePlayers(skill []int) int64 {
	slices.Sort(skill)
	l := 0
	r := len(skill) - 1
	teamSkill := 0
	var chemistry int64
	for l < r {
		if teamSkill != 0 && teamSkill != skill[l]+skill[r] {
			return -1
		}
		teamSkill = skill[l] + skill[r]
		chemistry += int64(skill[l] * skill[r])
		l++
		r--
	}
	return chemistry
}
