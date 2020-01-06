package easyslice

import (
	"fmt"
	"strconv"
)

func RemoveStringFromSlice(slc []string, s string) []string {
	newSlice := make([]string, 0)
	for _, v := range slc {
		if v != s {
			newSlice = append(newSlice, v)
		}
	}
	slc = nil
	return newSlice
}

func IterSlice(slc []string) {
	for _,v := range slc { fmt.Println(v) }
}

func IterSliceWithIndex(slc []string) {
	for i,v := range slc {
		fmt.Println("[" + strconv.Itoa(i) + "] " + v)
	}
}
