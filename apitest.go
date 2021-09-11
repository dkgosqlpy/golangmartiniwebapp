package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8088/courselist"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "localhost:8088")
	req.Header.Add("x-rapidapi-key", "testcodekey")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
