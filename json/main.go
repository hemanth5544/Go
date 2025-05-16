package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` //this will not be included in the json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json jinxx")

	// encodeJson()
	DecodeJson()

}

func encodeJson() {
	hemucor := []course{
		{"React JS", 299, "Udemy", "1234", []string{"web", "dev"}},
		{"Angular JS", 199, "Udemy", "1234", []string{"web", "dev"}},
	}

	fmt.Println("Hemanth course is", hemucor)

	jsonData, err := json.MarshalIndent(hemucor, "", "\t")

	if err != nil {
		fmt.Println("Error in encoding json", err)
		panic(err)
	}
	fmt.Println("Json data is", string(jsonData))
}

func DecodeJson() {
	jsonFromWeb := []byte(`
		{
			"name": "React JS",
			"price": 299,
			"platform": "Udemy",
			"tags": ["web", "dev"]
			}
			
			`)

	var core course

	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonFromWeb, &core)
		fmt.Printf("Course name is %s and price is %d", core.Name, core.Price)

	} else {
		fmt.Println("Json is not valid")
	}

}
