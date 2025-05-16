package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("srever jinxx")
	// GetRequest()
	PostRequest()
	// PerformFormPost()

}

func GetRequest() {
	const myurl = "https://monitra.webcookoo.com/api/wcse/2?url=https://datopic.ai"

	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println("Error in getting response", err)
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response satus code", response.StatusCode)

	fmt.Println("length of response", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println("my conten in stirng", string(content))
}

func PostRequest() {
	url := "https://api.freeapi.app/api/v1/kitchen-sink/http-methods/post"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func PerformFormPost() {
	url := "https://api.freeapi.app/api/v1/kitchen-sink/http-methods/post"

	payload := strings.NewReader("name=John&age=30")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
