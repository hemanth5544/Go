package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("srever jinxx")
	GetRequest()

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


func PostRequest(){
	
}