package main

import (
	"strconv"
)

func maximumSwap(num int) int {
	numStr := []byte(strconv.Itoa(num))
	nlen := len(numStr)
	maxIdx, idx0, idx1 := -1, -1, -1
	for i := nlen - 1; i >= 0; i-- {
		if maxIdx < 0 || numStr[i] > numStr[maxIdx] {
			maxIdx = i
		} else if numStr[i] < numStr[maxIdx] {
			idx0 = i
			idx1 = maxIdx
		}
	}
	if idx0 >= 0 {
		numStr[idx0], numStr[idx1] = numStr[idx1], numStr[idx0]
	}
	resNum, err := strconv.Atoi(string(numStr))
	if err != nil {
		panic(err)
	}
	return resNum
}
