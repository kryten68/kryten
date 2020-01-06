package EasyFiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileIntoSlice(p string) []string {

	fileBytes,err := ioutil.ReadFile(p)
		if err != nil { fmt.Println("Exception reading file: " + err.Error())}

	sliceOfLines := strings.Split(string(fileBytes),"\n")
	return sliceOfLines

}

func WriteSliceOfStringsIntoFile(d []string, p string) {

	f,err := os.OpenFile(p, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil { fmt.Println("Exception writing file: " + err.Error())}

	for _,v := range d {
		f.WriteString(v)
	}
	f.Close()

}