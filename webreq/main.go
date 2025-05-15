package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myurl = "https://monitra.webcookoo.com/api/wcse/2?url=https://datopic.ai"

func main() {

	fmt.Println("webreq jinxx")

	fmt.Println("myurl is", myurl)

	result ,_:=url.Parse(myurl)


	// Parse the URL

	fmt.Println("Parsed URL is", result)

	fmt.Println("Scheme is", result.Scheme)
	fmt.Println("Host is", result.Host)
	fmt.Println("Path is", result.Path)
	fmt.Println("Raw Query is", result.RawQuery)
	fmt.Println("Raw Path is", result.RawPath)
	

	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println("Error in getting response", err)
		panic(err)
	}

	fmt.Printf("Response Type :%T\n", response)

	fmt.Println("Response  bofyy ", response.Body)
	defer response.Body.Close() // Close the response body when done cox the go doesnt do it automatically

	data, err := io.ReadAll(response.Body) // Read the response body

	if err != nil {
		fmt.Println("Error in reading response", err)
		panic(err)
	}
	fmt.Println("Response Data", string(data)) // Convert the byte array to string and print it
}
