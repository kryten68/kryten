package easyslice

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
