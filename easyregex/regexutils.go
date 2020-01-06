package Easyregex

import (
	"regexp"
)

func GetSingleMatch(p string, s string) string {
	r,_ := regexp.Compile(p)
	v := r.FindAllString(s,-1)
	return v[0]
}

func GetMultipleMatch(p string, s string) []string {
	r,_ := regexp.Compile(p)
	v := r.FindAllString(s,-1)
	return v
}
