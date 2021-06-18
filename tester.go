package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func testmain() {

	url := "http://localhost:3333/api/videos"
	method := "GET"

	payload := strings.NewReader(`{` + "" + `"Title":"Hello",` + "" + `"Description":"Description",` + "" + `"Url": "url"` + "" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
