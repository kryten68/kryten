package Easyhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Get(url string) map[string]string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: "+ err.Error())
	}
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	responseMap := make(map[string]string)
	responseMap["response-code"] = strconv.Itoa(response.StatusCode)
	responseMap["response-status"] = response.Status
	responseMap["response-body"] = string(body)
	return responseMap
}