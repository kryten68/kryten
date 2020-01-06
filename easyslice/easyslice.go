package easyslice

import (
	"fmt"
)

func RemoveStringFromSlice(slc []string, s string) []string {

	newSlice := make([]string, 0)

	fmt.Println("Received: ")
	for _, v := range slc {
		fmt.Println(v)
	}

	for _, v := range slc {
		if v != s {
			newSlice = append(newSlice, v)
		}
	}

	slc = nil
	return newSlice

}
